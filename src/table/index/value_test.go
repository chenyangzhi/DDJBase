package index

import (
	"fmt"
	"os"
	"testing"
)

func TestHumanReadData(t *testing.T) {
	b := make([]byte, 2399) //设置读取的字节数
	f, _ := os.Open("/home/chenyangzhi/workplace/client_server/data/test/test/data")
	n, _ := f.Read(b)
	fmt.Println("total read ", n, " bytes!")
	iv := BytesToInternalValue(b[2277:2300])
	fmt.Println("the iv is", iv.IsRemove, iv.Key, iv.Length, iv.Value)
}
