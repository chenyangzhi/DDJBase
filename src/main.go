package main

import (
	"flag"
	"fmt"
	"os"
	"table/index"
	"time"
	logger "until/xlog4go"
)

var (
	size    = flag.Int("size", 1000000, "size of the tree to build")
	logconf = flag.String("l", "./conf/log.json", "log config file path")
)

func all(t *index.BTree) (out []index.Item) {
	t.Ascend(func(a index.Item) bool {
		out = append(out, a)
		return true
	})
	return
}

func GetIncreasedArray(size int) []int {
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		vals[i] = i
	}
	return vals
}

func main() {
	flag.Parse()
	if err := logger.SetupLogWithConf(*logconf); err != nil {
		panic(err)
	}
	defer logger.Close()
	table := index.NewTable("/home/chenyangzhi/workplace/client_server/data", "test", "test", "primaryKey")
	f := table.CreateTable()
	//vals := rand.Perm(*size)
	vals := GetIncreasedArray(*size)
	tr := index.BuildBTreeFromPage(table.GetTablePath(), f)
	t := time.Now()
	for _, v := range vals {
		bs := "13434534534adsfgdsafffffffffffffgagagsadgsdfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffretsrgsdfgdsfgsdfgdf"
		tr.Insert(uint64(v), []byte(bs))
	}

	elapsed := time.Since(t)
	fmt.Println("the insert 1000000 key time elapsed ", elapsed)
	//t = time.Now()
	count := 0
	for _, v := range vals {
		val := tr.GetByKey(uint64(v))
		vv := string(val)
		if val != nil {
			if v == 999999 {
				fmt.Println("the key is ", v, " the value is", vv)
			}
		} else {
			count++
			fmt.Println("error: not found val = ", v)
		}
	}
	elapsed = time.Since(t)
	fmt.Println("the time elapsed ", elapsed)
	//root := tr.GetRootNode()
	//root.Print(os.Stdout, 2)
	fmt.Println("the not fount count is ", count)
	fmt.Println("the time elapsed ", elapsed)
	tr.Commit()
	os.Exit(0)

}
