package main

import (
	"errors"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "kvprobuf"
	"time"
	logger "until/xlog4go"
)

const (
	address = "localhost:50051"
)

var (
	c       pb.DDJkvClient
	logConf = flag.String("l", "./conf/log.json", "log config file path")
)

type Iterator func() ([]byte, error)

func InitClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Error("did not connect: %v", err)
	}
	c = pb.NewDDJkvClient(conn)
}

func NewClosureIterator(b uint64) Iterator {
	l := []*pb.Slice{}
	var current *pb.Slice
	idx := b
	first := true
	fn := func() ([]byte, error) {
		if len(l) == 0 {
			r, err := c.AscendIterate(context.Background(), &pb.Slice{Key: idx, Value: []byte{}, IncludeStart: first})
			if err != nil || r == nil || len(r.Slice) == 0 {
				return nil, errors.New("StopIteration")
			}
			l = r.Slice
			first = false
		}
		current = l[0]
		l = l[1:]
		idx = current.Key
		val := current.Value
		return val, nil
	}
	return fn
}

func find(key uint64) Iterator {
	return NewClosureIterator(key)
}

func Put(key uint64, value []byte) {
	t := time.Now()
	logger.Info("the requeset start!")
	pb, err := c.Put(context.Background(), &pb.Slice{Key: key, Value: value})
	if err != nil || pb.Success != true {
		logger.Error("the the error is", err)
	}
	elapsed := time.Since(t)
	logger.Info("the requeset finished in %v", elapsed)
	return
}

func Get(key uint64) []byte {
	pb, err := c.Get(context.Background(), &pb.Slice{Key: key})
	if err != nil {
		logger.Error("the the error is", err)
		return nil
	}
	return pb.Slice.Value
}

func Delete(key uint64) []byte {
	pb, err := c.Delete(context.Background(), &pb.Slice{Key: key})
	if err != nil {
		logger.Error("the the error is", err)
		return nil
	}
	if pb.Success == true {
		return pb.Value
	}
	return nil
}

func main() {
	flag.Parse()
	if err := logger.SetupLogWithConf(*logConf); err != nil {
		panic(err)
	}
	defer logger.Close()
	InitClient()
	next := find(uint64(300))
	t := time.Now()
	for {
		_, err := next()
		if err != nil {
			break
		}
		//logger.Info("the value is %v", string(v))
	}
	elapsed := time.Since(t)
	logger.Info("Put 100000000 key in DDJkv elapsed %s", elapsed)
	//key := uint64(125)
	//t := time.Now()
	//for i := 0; i < 1000000; i++ {
	//	if i%1000 == 0 {
	//		logger.Info("%v key is put ", i)
	//	}
	//	str := "1234asdfsdfsdfsadfsdfsdfasfddfasdfasdfasdfaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaafsdfsdfsdfasdfadsfadfadsfadfaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaafasdffasdf"
	//	Put(uint64(i), []byte(str))
	//}
	//elapsed := time.Since(t)
	//logger.Info("Put 100000000 key in DDJkv elapsed %s", elapsed)
	//b := Get(key)
	//if len(b) != 0 {
	//	logger.Info("Get: the key is %v%s%v", key, " the value is ", string(b))
	//} else {
	//	logger.Info("Get: the key is %v %s", key, " is not found!")
	//}
	//b = Delete(key)
	//if b != nil {
	//	logger.Info("Delete: the key is %v%s%v", key, " the value is ", string(b))
	//}
}
