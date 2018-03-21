package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"runtime"
	"table/index"
	"time"
	logger "until/xlog4go"
	"container/list"
	"errors"
	"log"
)

var (
	size    = flag.Int("size", 100, "size of the tree to build")
	logconf = flag.String("l", "./conf/log.json", "log config file path")
)

func all(t *index.BTree) (out []index.Item) {
	t.Ascend(func(a index.Item) bool {
		out = append(out, a)
		return true
	})
	return
}

func memery() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		log.Printf("Alloc = %vMB  TotalAlloc = %vMB  Sys = %vMB  NumGC = %vMB\n", m.Alloc/(1024*1024), m.TotalAlloc/(1024*1024), m.Sys/(1024*1024), m.NumGC)
		time.Sleep(5 * time.Second)
	}
}

type Iterator func() (value *list.Element, err error)
func NewClosureIterator(b index.BtreeNodeItem,tr *index.BTree) Iterator{
	l := list.New()
	fmt.Println("the l length is ",l.Len())
	var current *list.Element = l.Front()
	var r *list.Element
	markItem := b
	fn := func() (value *list.Element,err error) {
		r = current
		if r == nil {
			l = AscendGreaterOrEqual(markItem,tr)
			current = l.Front()
			r =  current
			length := l.Len()
			if  length == 0{
				return nil, errors.New("StopIteration")
			}
		}
		markItem = *(current.Value.(*index.BtreeNodeItem))
		current = current.Next()
		return r, nil
	}
	return fn
}
func AscendGreaterOrEqual(b index.BtreeNodeItem,tr *index.BTree) *list.List{
	got := list.New()
	tr.AscendGreaterOrEqual(b, func(a index.Item) bool {
		got.PushBack(a)
		if got.Len() >= 100 {
			return false
		}
		return true
	})
	return got
}

func find(key uint64,tr *index.BTree) Iterator {
	var b index.BtreeNodeItem
	b.IdxId = key
	return NewClosureIterator(b,tr)
}

func GetIncreasedArray(size int) []int{
	vals := make([]int,size)
	for i := 0; i < size ; i++ {
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
	go memery()
	table := index.NewTable("/home/chenyangzhi/workplace/client_server/data", "test", "test", "primaryKey")
	table.CreateTable()
	//vals := rand.Perm(*size)
	vals := GetIncreasedArray(*size)
	tr := index.BuildBTreeFromPage(table.GetTablePath())
	t := time.Now()
	//for _, v := range vals {
	//	var b index.BtreeNodeItem
	//	bs := make([]byte,8,8)
	//	b.IdxId = uint64(v)
	//	binary.LittleEndian.PutUint64(bs, uint64(b.IdxId))
	//	b.Key = bs
	//	if b.IdxId == 178 {
	//		tr.ReplaceOrInsert(&b)
	//	}else{
	//		tr.ReplaceOrInsert(&b)
	//	}
	//	item := tr.Get(&b)
	//	if item == nil {
	//		fmt.Println("error: not insert val = ", v)
	//	}
	//}
	//elapsed := time.Since(t)
	//fmt.Println("the time elapsed ", elapsed)
	//t = time.Now()
	next := find(uint64(64),tr)
	for {
		v, err := next()
		if err != nil {
			break
		}
		val := v.Value.(*index.BtreeNodeItem).IdxId
		fmt.Println(val)
	}
	count := 0
	for _, v := range vals {
		var b index.BtreeNodeItem
		b.IdxId = uint64(v)
		bs := make([]byte, 8, 8)
		binary.LittleEndian.PutUint64(bs, uint64(b.IdxId))
		b.Key = bs
		item := (tr.Get(&b)).(*index.BtreeNodeItem)
		//fmt.Println("the key is ",item.GetBtreeNodeItemID(), " the value is",binary.LittleEndian.Uint64(item.GetBtreeNodeItemKey()))
		if item == nil {
			count++
			fmt.Println("error: not found val = ", v)
		}
	}
	elapsed := time.Since(t)
	//root := tr.GetRootNode()
	//root.Print(os.Stdout, 2)
	fmt.Println("the not fount count is ", count)
	fmt.Println("the time elapsed ", elapsed)
	fmt.Println("the tree all of node id ", tr.GetNodeIds())
	set := tr.GetDirtyPage()
	fmt.Println("the dirty page is %v ", set)
	tr.Commit()
	fmt.Println("the dirty page is commit")
	os.Exit(0)

}
