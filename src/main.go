package main

import (
	"flag"
	"fmt"
	"table/index"
	"time"
	logger "until/xlog4go"
)

var (
	size    = flag.Int("size", 10000000, "size of the tree to build")
	logconf = flag.String("l", "./conf/log.json", "log config file path")
	vals    []int
)

func GetIncreasedArray(size int) []int {
	vals := make([]int, size)
	for i := 0; i < size; i++ {
		vals[i] = i
	}
	return vals
}

func TestInsert() {
	table := index.NewTable("./data", "test", "test")
	f := table.CreateNewTable()
	//vals := rand.Perm(*size)
	index.Curtr = index.BuildBTreeFromPage(table.GetIndexPath(), f)
	ts := int64(0)
	maxts := int64(0)
	count := 0
	for _, v := range vals {
		ts = time.Now().UnixNano()
		bs := "13434567890"
		index.Insert(uint64(v), []byte(bs))
		tmp := time.Now().UnixNano() - ts
		if tmp > maxts {
			maxts = tmp
		}
		if tmp > 10000000 {
			count++
		}

	}
	fmt.Println("the max put elapsed:", maxts)
	fmt.Println("the more thar 10ms count:", count)
}

func TestGet() {
	table := index.NewTable("./data", "test", "test")
	f := table.CreateNewTable()
	//vals := rand.Perm(*size)
	index.Curtr = index.BuildBTreeFromPage(table.GetIndexPath(), f)
	count := 0
	ts := int64(0)
	maxts := int64(0)
	than := 0
	for _, v := range vals {
		ts = time.Now().UnixNano()
		val := index.Curtr.GetByKey(uint64(v))
		if val == nil || len(val) == 0 {
			fmt.Println("error: not found val = ", v)
			count++
		}
		tmp := time.Now().UnixNano() - ts
		if tmp > maxts {
			maxts = tmp
		}
		if tmp > 10000000 {
			than++
		}
	}
	fmt.Println("not found key has ", count)
	fmt.Println("the max put elapsed:", maxts)
	fmt.Println("the get more than 10ms count:", than)
}
func TestInsertAndGet() {
	table := index.NewTable("./data", "test", "test")
	f := table.CreateNewTable()
	//vals := rand.Perm(*size)
	index.Curtr = index.BuildBTreeFromPage(table.GetIndexPath(), f)
	t := time.Now()
	for _, v := range vals {
		bs := "13434567890"
		index.Insert(uint64(v), []byte(bs))
	}
	elapsed := time.Since(t)
	logger.Info("the insert 10000000 key time elapsed ", elapsed)
	for _, v := range vals {
		val := index.Curtr.GetByKey(uint64(v))
		if val == nil || len(val) == 0 {
			fmt.Println("error: not found val = ", v)
		}
	}
	elapsed = time.Since(t)
	logger.Info("the get 10000000 key time elapsed ", elapsed)
}

func TestDelete() {

}

func TestIterate() {
	var bItem index.BtreeNodeItem
	table := index.NewTable("./data", "test", "test")
	f := table.CreateNewTable()
	//vals := rand.Perm(*size)
	index.Curtr = index.BuildBTreeFromPage(table.GetIndexPath(), f)
	for i := 0; i < 10000; i++ {
		t := time.Now()
		bItem.IdxId = uint64(i*1000 + 1000)
		index.AscendIterate(bItem, index.Curtr, true)
		elapsed := time.Since(t)
		logger.Info("the iterate 1000 key time elapsed ", elapsed)
	}
}

func main() {
	flag.Parse()
	if err := logger.SetupLogWithConf(*logconf); err != nil {
		panic(err)
	}
	defer logger.Close()
	vals = GetIncreasedArray(*size)
	t := time.Now()
	TestInsert()
	//TestGet()
	//TestInsertAndGet()
	//TestIterate()
	elapsed := time.Since(t)
	logger.Info("the time elapsed ", elapsed)
	for {
		time.Sleep(1 * time.Second)
		index.Curtr.Commit()
	}
}
