package index
import (
	"os"
	"log"
	"common"
	"encoding/binary"
	"sort"
	"container/list"
)

const(
	MAXPOOLSIZE = 24
)
type ValuePool struct {
	EofOffset        uint64
	CurrentOffest    uint64
	RemoveValues    *list.List
	MemeryTail      []byte
}

func NewValuePool(eof uint64) *ValuePool{
	rlist := list.New()
	b := make([]byte,MAXPOOLSIZE)
	return &ValuePool{
		EofOffset: eof,
		CurrentOffest: 0,
		RemoveValues: rlist,
		MemeryTail:   b,
	}
}

type InternalValue struct {
	IsRemove  bool
	Length    uint32
	Key       uint64
	Value     []byte
}

func (iv1 InternalValue) Less(iv2 InternalValue) bool {
	return iv1.Key < iv2.Key
}
func NewInternalValue(isRemove bool,length uint32, key uint64, val []byte) *InternalValue{
	return &InternalValue{
		IsRemove:  isRemove,
		Length:    length,
		Key:       key,
		Value:     val,
	}

}

func (iv InternalValue) Size() uint32{
	return uint32(common.INT8_LEN + common.INT64_LEN + common.INT32_LEN + len(iv.Value))
}

func(iv InternalValue) ToBytes()[]byte{
	b := make([]byte,iv.Size() + CRCSIZE)
	iStart, iEnd := 0, 0
	if iv.IsRemove == true{
		b[iStart] = 1
	}else{
		b[iStart] = 0
	}
	iStart += common.INT8_LEN
	iEnd = iStart + common.INT32_LEN
	binary.LittleEndian.PutUint32(b[iStart:iEnd], iv.Length)
	iStart = iEnd
	iEnd = iStart + common.INT64_LEN
	binary.LittleEndian.PutUint64(b[iStart:iEnd], iv.Key)
	iStart = iEnd
	iEnd = iStart + len(iv.Value)
	copy(b[iStart:iEnd],iv.Value)
	crc := common.Crc16(b[0:iEnd])
	iStart = iEnd
	iEnd = iStart + CRCSIZE
	binary.LittleEndian.PutUint16(b[iStart:iEnd], crc)
	return b
}

func BytesToInternalValue(b []byte) *InternalValue{
	iStart , iEnd := 0, 0
	remove := false
	if b[0] != 0 {
		remove = true
	}
	iStart = iStart + 1
	iEnd = iStart + common.INT32_LEN
	length := binary.LittleEndian.Uint32(b[iStart:iEnd])
	iStart = iEnd
	iEnd = iStart + common.INT64_LEN
	key := binary.LittleEndian.Uint64(b[iStart:iEnd])
	value := make([]byte,length)
	if remove == true{
		return NewInternalValue(true,0,key,[]byte{})
	}
	iStart = iEnd
	iEnd = iStart + int(length)
	copy(value,b[iStart:iEnd])
	crc_0 := common.Crc16(b[0:iEnd])
	iStart = iEnd
	iEnd = iStart + common.INT16_LEN
	crc_1 := binary.LittleEndian.Uint16(b[iStart:iEnd])
	_assert(crc_0 == crc_1, "the InternalValue crc is failed")
	return NewInternalValue(false,length,key,value)
}

func GetFileOffset(f *os.File) uint64{
	currentPosition, _ := f.Seek(0, 2)
	return uint64(currentPosition)
}

func WriteAt(whence uint64,w []byte,f *os.File){
	if _,err :=f.WriteAt(w,int64(whence)); err != nil {
		log.Fatal(err)
	}
}

func ReadAt(offset uint64,cow *copyOnWriteContext)*InternalValue{
	if offset >= cow.vPool.EofOffset {
		return MemReadAt(offset,cow)
	}else{
		return FileReadAt(offset,cow.dataFile)
	}
}

func BatchReadAt(offsets []uint64,cow *copyOnWriteContext) []*InternalValue{
	sort.Slice(offsets, func(i, j int) bool { return offsets[i] < offsets[j] })
	memOffset := make([]uint64,0,2)
	fileOffset := make([]uint64,0,2)
	values := make([]*InternalValue,0,2)
	for _,off := range offsets{
		if off >= cow.vPool.EofOffset {
			memOffset = append(memOffset,off)
		}else{
			fileOffset = append(fileOffset,off)
		}
	}
	v1 := BatchMemReadAt(memOffset,cow)
	v2 := BatchFileReadAt(fileOffset,cow.dataFile)
	values = append(values,v1...)
	values = append(values,v2...)
	sort.Slice(values, func(i, j int) bool { return values[i].Key < values[j].Key })
	return values
}
func MemReadAt(off uint64,cow *copyOnWriteContext) *InternalValue{
	relativeOff := off - cow.vPool.EofOffset
	length := binary.LittleEndian.Uint32(cow.vPool.MemeryTail[relativeOff+1:relativeOff + 5])
	iLen := length + common.INT8_LEN + common.INT64_LEN + common.INT32_LEN + CRCSIZE
	bs := make([]byte,iLen)
	copy(bs,cow.vPool.MemeryTail[relativeOff: relativeOff + uint64(iLen)])
	return BytesToInternalValue(bs)
}
func BatchMemReadAt(offsets []uint64,cow *copyOnWriteContext) []*InternalValue{
	values := make([]*InternalValue,len(offsets))
	for i,off := range offsets{
		iv := MemReadAt(off,cow)
		values[i] = iv
	}
	return values
}

func FileReadAt(off uint64,f *os.File)*InternalValue{
	lbyte := make([]byte,common.INT32_LEN)
	if _,err := f.ReadAt(lbyte,int64(off) + int64(common.INT8_LEN)); err != nil{
		log.Fatal(err)
	}
	length := binary.LittleEndian.Uint32(lbyte)
	bs := make([]byte,length + common.INT8_LEN + common.INT64_LEN + common.INT32_LEN + CRCSIZE)
	if _,err := f.ReadAt(bs,int64(off)); err != nil{
		log.Fatal(err)
	}
	return BytesToInternalValue(bs)
}

func BatchFileReadAt(offsets []uint64,f *os.File) []*InternalValue{
	values := make([]*InternalValue,len(offsets))
	for i,off := range offsets{
		iv := FileReadAt(off,f)
		values[i] = iv
	}
	return values
}
