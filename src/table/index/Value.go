package index
import (
	"os"
	"log"
	"common"
	"encoding/binary"
)
type InternalValue struct {
	IsRemove  bool
	Length    uint32
	Value     []byte
}

func CreateDataFile(filePath string)*os.File{
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func GetFileOffset(f *os.File) uint64{
	currentPosition, _ := f.Seek(0, 2)
	return currentPosition
}

func Append(values []InternalValue,f *os.File){
	totalLength := 0
	for _,v := range values {
		totalLength += v.Length + 1 + 4
	}
	barr := make([]byte,totalLength)
	iStart := 0
	for _,v := range values{
		barr[iStart] = byte(v.IsRemove)
		iStart += iStart + 1
		binary.LittleEndian.PutUint32(barr[iStart:iStart+4], v.Length)
		iStart += iStart + 4
		copy(barr[iStart:iStart+v.Length],v.Value)
	}
	if _, err := f.Write(barr); err != nil {
		log.Fatal(err)
	}
}

func WriteAt(offset uint64,w []byte,f *os.File){
	if _,err :=f.WriteAt(w,offset); err != nil {
		log.Fatal(err)
	}
}

func ReadAt(offsets []uint64,f *os.File) []InternalValue{
	values := make([]*InternalValue,len(offsets))
	lbyte := make([]byte,common.INT32_LEN + 1)
	for _,off := range offsets{
		if _,err := f.ReadAt(lbyte,off); err != nil{
			log.Fatal(err)
		}
		length := binary.LittleEndian.Uint32(lbyte[1:])
		value := make([]byte,length)
		if _,err := f.ReadAt(lbyte,value); err != nil{
			log.Fatal(err)
		}
		iv := &InternalValue{
			IsRemove:  true,
			Length:    length,
			Value:     value,
		}
		values = append(values,iv)
	}
	return values
}
