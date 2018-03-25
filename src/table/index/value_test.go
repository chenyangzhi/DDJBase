package index

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"until"
)

func TestHumanReadData(t *testing.T) {
	fmt.Println(until.Goid())
	path := "/some/path/to/remove/file.name"
	file := filepath.Base(path)
	fmt.Println(file)
	b := make([]byte, 5100) //设置读取的字节数
	f, _ := os.OpenFile("/home/chenyangzhi/workplace/client_server/data/test/test/data_1", os.O_RDWR, os.ModePerm)
	n, _ := f.Read(b)
	fmt.Println("total read ", n, " bytes!")
	iStart := 0
	iEnd := iStart + 255
	for iEnd <= 5100 {
		iv := BytesToInternalValue(b[iStart:iEnd])
		iStart = iEnd
		iEnd = iStart + 255
		fmt.Println("the iv is", iv.IsRemove, iv.Key, iv.Length, string(iv.Value))
	}
}
