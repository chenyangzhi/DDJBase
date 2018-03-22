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
	size    = flag.Int("size", 300, "size of the tree to build")
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
	var current *list.Element = l.Front()
	var r *list.Element
	markItem := b
	first := true
	fn := func() (value *list.Element,err error) {
		r = current
		if r == nil {
			l = index.AscendGreaterOrEqual(markItem,tr,first)
			first = false
			current = l.Front()
			r =  current
			length := l.Len()
			if  length == 0{
				return nil, errors.New("StopIteration")
			}
		}
		idx := current.Value.(*index.InternalValue).Key
		markItem.IdxId = idx
		current = current.Next()
		return r, nil
	}
	return fn
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
	f := table.CreateTable()
	//vals := rand.Perm(*size)
	vals := GetIncreasedArray(*size)
	tr := index.BuildBTreeFromPage(table.GetTablePath(),f)
	t := time.Now()
	for _, v := range vals {
		bs := make([]byte,8)
		binary.LittleEndian.PutUint64(bs,uint64(v))
		tr.Insert(uint64(v),bs)
	}

	elapsed := time.Since(t)
	fmt.Println("the time elapsed ", elapsed)
	//t = time.Now()
	next := find(uint64(64),tr)
	for {
		v, err := next()
		if err != nil {
			break
		}
		val := binary.LittleEndian.Uint64(v.Value.(*index.InternalValue).Value)
		fmt.Println("the value is ",val)
	}
	elapsed = time.Since(t)
	fmt.Println("the time elapsed ", elapsed)
	count := 0
	for _, v := range vals {
		val := tr.GetByKey(uint64(v))
		vv := binary.LittleEndian.Uint64(val)
		if val != nil {
			if vv == 99 {
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
