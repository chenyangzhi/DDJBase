package index

import (
	"fmt"
	"os"
	"testing"
)

func TestHumanReadData(t *testing.T) {
	length := 265
	offset := 3295 * 264
	b := make([]byte, length) //设置读取的字节数
	f, _ := os.OpenFile("/home/chenyangzhi/workplace/client_server/data/test/test/data_1522033120", os.O_RDWR, os.ModePerm)
	n, _ := f.ReadAt(b, int64(offset))
	fmt.Println("total read ", n, " bytes!")
	iStart := 0
	iEnd := iStart + 265
	count := 0
	i := uint64(0)
	for iEnd <= length {
		iv := BytesToInternalValue(b[iStart:iEnd])
		fmt.Println(b[iStart:iEnd])
		iStart = iEnd
		iEnd = iStart + 265
		if iv.IsRemove == true {
			fmt.Println("the iv is", iv.IsRemove, iv.Key, iv.Length, string(iv.Value))
			count++
		}
		//if iv.Key != i {
		//	fmt.Println("not found the key",i)
		//}
		i++
		fmt.Println("the iv is", iv.IsRemove, iv.Key, iv.Length, string(iv.Value))
	}
	fmt.Println("the count is", count)
}
