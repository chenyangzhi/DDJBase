package main

import (
	"flag"
	"fmt"
	"table/index"
	"time"
	logger "until/xlog4go"
)

var (
	size    = flag.Int("size", 100000000, "size of the tree to build")
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
	table := index.NewTable("./data", "test", "test")
	f := table.CreateNewTable()
	//vals := rand.Perm(*size)
	vals := GetIncreasedArray(*size)
	index.Curtr = index.BuildBTreeFromPage(table.GetIndexPath(), f)
	t := time.Now()
	for _, v := range vals {
		bs := "13434534534adsfgdsafffffffffffffgagagsadgsdfffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffretsrgsdfgdsfgsdfgdf"
		index.Insert(uint64(v), []byte(bs))
	}

	elapsed := time.Since(t)
	fmt.Println("the insert 1000000 key time elapsed ", elapsed)
	//t = time.Now()
	count := 0
	for _, v := range vals {
		val := index.Curtr.GetByKey(uint64(v))
		if val == nil || len(val) == 0 {
			fmt.Println("error: not found val = ", v)
		}
	}
	elapsed = time.Since(t)
	fmt.Println("the time elapsed ", elapsed)
	//root := tr.GetRootNode()
	//root.Print(os.Stdout, 2)
	fmt.Println("the not fount count is ", count)
	fmt.Println("the time elapsed ", elapsed)
	for {
		time.Sleep(1 * time.Second)
		index.Curtr.Commit()
	}
}
