package index

import (
	"common"
	"container/list"
	"encoding/binary"
	"os"
	"sort"
	"sync"
	logger "until/xlog4go"
)

const (
	MAXPOOLSIZE   = 1024 * 1024 * 2
	MEMQUEUESIZE  = 10
	VACCUMSIZE    = 1024 * 1024 * 2
	GCMINFILESIZE = 1024 * 1024 * 64
)

var (
	Curtr        *BTree
	VaccumHolder *VaccumData
)

type ValuePool struct {
	EofOffset     uint64
	CurrentOffest uint64
	RemoveValues  *list.List
	MemeryTail    []byte
	MtMutex       *sync.Mutex
}

type VaccumData struct {
	VaccumFlag bool
	VMutex     *sync.Mutex
}

func InitValuePoolQueue(size int) chan *ValuePool {
	c := make(chan *ValuePool, size)
	for i := 0; i < size; i++ {
		vp := NewValuePool(0)
		c <- vp
	}
	return c
}

func NewValuePool(eof uint64) *ValuePool {
	rlist := list.New()
	b := make([]byte, MAXPOOLSIZE)
	return &ValuePool{
		EofOffset:     eof,
		CurrentOffest: 0,
		RemoveValues:  rlist,
		MemeryTail:    b,
		MtMutex:       new(sync.Mutex),
	}
}

type InternalValue struct {
	IsRemove bool
	Length   uint32
	Key      uint64
	Value    []byte
}

func (iv1 InternalValue) Less(iv2 InternalValue) bool {
	return iv1.Key < iv2.Key
}
func NewInternalValue(isRemove bool, length uint32, key uint64, val []byte) *InternalValue {
	return &InternalValue{
		IsRemove: isRemove,
		Length:   length,
		Key:      key,
		Value:    val,
	}

}

func (iv InternalValue) Size() uint32 {
	return uint32(common.INT8_LEN + common.INT64_LEN + common.INT32_LEN + len(iv.Value))
}

func (iv InternalValue) ToBytes() []byte {
	b := make([]byte, iv.Size()+CRCSIZE)
	iStart, iEnd := 0, 0
	if iv.IsRemove == true {
		b[iStart] = 1
	} else {
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
	copy(b[iStart:iEnd], iv.Value)
	crc := common.Crc16(b[0:iEnd])
	iStart = iEnd
	iEnd = iStart + CRCSIZE
	binary.LittleEndian.PutUint16(b[iStart:iEnd], crc)
	return b
}

func BytesToInternalValue(b []byte) *InternalValue {
	iStart, iEnd := 0, 0
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
	value := make([]byte, length)
	if remove == true {
		return NewInternalValue(true, 0, key, []byte{})
	}
	iStart = iEnd
	iEnd = iStart + int(length)
	copy(value, b[iStart:iEnd])
	crc_0 := common.Crc16(b[0:iEnd])
	iStart = iEnd
	iEnd = iStart + common.INT16_LEN
	crc_1 := binary.LittleEndian.Uint16(b[iStart:iEnd])
	_assert(crc_0 == crc_1, "the InternalValue crc is failed")
	return NewInternalValue(false, length, key, value)
}

func GetFileOffset(f *os.File) uint64 {
	currentPosition, _ := f.Seek(0, 2)
	return uint64(currentPosition)
}

func WriteAt(whence uint64, w []byte, f *os.File) {
	_, err := f.WriteAt(w, int64(whence))
	if err != nil {
		logger.Error("WriteAt is error: %v", err)
	}
}

func ReadAt(offset uint64, cow *copyOnWriteContext) *InternalValue {
	if offset >= cow.curFileOffset {
		return MemReadAt(offset, cow)
	} else {
		return FileReadAt(offset, cow.dataFile)
	}
}

func BatchReadAt(offsets []uint64, cow *copyOnWriteContext) []*InternalValue {
	sort.Slice(offsets, func(i, j int) bool { return offsets[i] < offsets[j] })
	memOffset := make([]uint64, 0, 2)
	fileOffset := make([]uint64, 0, 2)
	values := make([]*InternalValue, 0, 2)
	for _, off := range offsets {
		if off >= cow.curVPool.EofOffset {
			memOffset = append(memOffset, off)
		} else {
			fileOffset = append(fileOffset, off)
		}
	}
	v1 := BatchMemReadAt(memOffset, cow)
	v2 := BatchFileReadAt(fileOffset, cow.dataFile)
	values = append(values, v1...)
	values = append(values, v2...)
	sort.Slice(values, func(i, j int) bool { return values[i].Key < values[j].Key })
	return values
}

func ValuePoolReadAt(vPool *ValuePool, off uint64) *InternalValue {
	vPool.MtMutex.Lock()
	if off < vPool.EofOffset || off >= vPool.EofOffset+vPool.CurrentOffest {
		vPool.MtMutex.Unlock()
		return nil
	}
	relativeOff := off - vPool.EofOffset
	length := binary.LittleEndian.Uint32(vPool.MemeryTail[relativeOff+1 : relativeOff+5])
	iLen := length + common.INT8_LEN + common.INT64_LEN + common.INT32_LEN + CRCSIZE
	bs := make([]byte, iLen)
	copy(bs, vPool.MemeryTail[relativeOff:relativeOff+uint64(iLen)])
	vPool.MtMutex.Unlock()
	return BytesToInternalValue(bs)
}

func MemReadAt(off uint64, cow *copyOnWriteContext) *InternalValue {
	cur := cow.curVPool
	iv := ValuePoolReadAt(cur, off)
	if iv != nil {
		return iv
	}
	l := cow.fullQueue
	for e := l.Front(); e != nil; e = e.Next() {
		iv = ValuePoolReadAt(e.Value.(*ValuePool), off)
		if iv != nil {
			return iv
		}
	}
	l = cow.flushQueue
	for e := l.Front(); e != nil; e = e.Next() {
		iv = ValuePoolReadAt(e.Value.(*ValuePool), off)
		if iv != nil {
			return iv
		}
	}
	return iv
}
func BatchMemReadAt(offsets []uint64, cow *copyOnWriteContext) []*InternalValue {
	values := make([]*InternalValue, len(offsets))
	for i, off := range offsets {
		iv := MemReadAt(off, cow)
		values[i] = iv
	}
	return values
}

func FileReadAt(off uint64, f *os.File) *InternalValue {
	lbyte := make([]byte, common.INT32_LEN)
	if _, err := f.ReadAt(lbyte, int64(off)+int64(common.INT8_LEN)); err != nil {
		logger.Error("FileReadAt is error: %v", err)
	}
	length := binary.LittleEndian.Uint32(lbyte)
	bs := make([]byte, length+common.INT8_LEN+common.INT64_LEN+common.INT32_LEN+CRCSIZE)
	if _, err := f.ReadAt(bs, int64(off)); err != nil {
		logger.Error("FileReadAt is error: %v", err)
	}
	return BytesToInternalValue(bs)
}

func BatchFileReadAt(offsets []uint64, f *os.File) []*InternalValue {
	values := make([]*InternalValue, len(offsets))
	for i, off := range offsets {
		iv := FileReadAt(off, f)
		values[i] = iv
	}
	return values
}

func Flush(cow *copyOnWriteContext) {
	var vp *ValuePool
	for e := cow.flushQueue.Front(); e != nil; e = e.Next() {
		vp = e.Value.(*ValuePool)
		WriteAt(vp.EofOffset, vp.MemeryTail[0:vp.CurrentOffest], cow.dataFile)
		cow.curFileOffset = vp.EofOffset + vp.CurrentOffest
		vp.CurrentOffest = 0
		aList := vp.RemoveValues
		for e := aList.Front(); e != nil; e = e.Next() {
			off := e.Value.(*BtreeNodeItem).Key
			WriteAt(off, []byte{1}, cow.dataFile)
		}
		vp.RemoveValues = list.New()
		cow.emptyQueue <- vp
	}
	cow.flushQueue = list.New()
	cow.flushFinish = true
}
func BatchBytesToInternalValues(bs []byte) ([]*InternalValue, uint64) {
	iStart, iEnd := 0, 0
	iv := make([]*InternalValue, 0, 16)
	l := len(bs)
	for {
		if iStart+5 < l {
			length := binary.LittleEndian.Uint32(bs[iStart+1 : iStart+5])
			totalSz := length + common.INT8_LEN + common.INT64_LEN + common.INT32_LEN + CRCSIZE
			iEnd = iStart + int(totalSz)
			if iEnd <= l {
				v := BytesToInternalValue(bs[iStart:iEnd])
				iv = append(iv, v)
				iStart = iEnd
			} else {
				break
			}

		} else {
			break
		}

	}
	return iv, uint64(iEnd)
}

func VaccumInsert(bs []byte, off uint64, f *os.File, tr *BTree) uint64 {
	_, err := f.ReadAt(bs, int64(off))
	if err != nil {
		logger.Info("read file is error: %v", err)
	}
	iv, n := BatchBytesToInternalValues(bs)
	off = off + n
	for _, o := range iv {
		if o.IsRemove == false {
			tr.InternalInsert(o.Key, o.Value)
		}
	}
	return off
}

func Vaccum(cow *copyOnWriteContext) {
	logger.Info("this time vaccum starting")
	table := NewTable("./data", "test", "test", "primaryKey_1", "data_1")
	f := table.CreateTable()
	btr := BuildBTreeFromPage(table.GetTablePath(), f)
	dataFile := cow.dataFile
	sentryOffset := cow.curFileOffset
	iStart, iEnd := uint64(0), uint64(0)
	bs := make([]byte, VACCUMSIZE)
	for iEnd < sentryOffset {
		iStart = VaccumInsert(bs, iStart, dataFile, btr)
		sentryOffset = cow.curFileOffset
		iEnd = iStart + VACCUMSIZE
	}
	iStart = VaccumInsert(bs[0:(sentryOffset-iStart)], iStart, dataFile, btr)
	VaccumHolder.VMutex.Lock()
	fileOffset := cow.curVPool.EofOffset + cow.curVPool.CurrentOffest
	cow.fullQueue.PushBack(cow.curVPool)
	cow.flushQueue = cow.fullQueue
	cow.fullQueue = list.New()
	Flush(cow)
	sentryOffset = fileOffset
	l := sentryOffset - iStart
	bs = make([]byte, l)
	VaccumInsert(bs[0:l], iStart, dataFile, btr)
	tmp := Curtr
	Curtr = btr
	VaccumHolder.VaccumFlag = false
	VaccumHolder.VMutex.Unlock()
	freeBtree(tmp)
	logger.Info("this time vaccum finished")
}
