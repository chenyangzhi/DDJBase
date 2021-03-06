package main

import (
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "kvprobuf"
	"net"
	"table/index"
	"until"
	logger "until/xlog4go"
)

var (
	logConf = flag.String("l", "./conf/log.json", "log config file path")
)

const (
	port = ":50051"
)

type server struct{}

var (
	tr *index.BTree
)

func (s *server) AscendIterate(ctx context.Context, param *pb.Slice) (*pb.AscendIterateResponse, error) {
	var bItem index.BtreeNodeItem
	bItem.IdxId = param.Key
	b := new(pb.AscendIterateResponse)
	l := index.AscendIterate(bItem, tr, param.IncludeStart)
	b.Slice = make([]*pb.Slice, 0, 32)
	for _, iv := range l {
		key := iv.Key
		value := iv.Value
		t := &pb.Slice{
			Key:   key,
			Value: value,
		}
		b.Slice = append(b.Slice, t)
	}
	return b, nil
}

func (s *server) Put(ctx context.Context, param *pb.Slice) (*pb.PutResponse, error) {
	var p pb.PutResponse
	index.Insert(param.Key, param.Value)
	p.Success = true
	return &p, nil
}

func (s *server) Get(ctx context.Context, param *pb.Slice) (*pb.GetResponse, error) {
	var p pb.GetResponse
	var slice pb.Slice
	val := tr.GetByKey(param.Key)
	slice.Value = val
	slice.Key = param.Key
	p.Slice = &slice
	return &p, nil
}

func (s *server) Delete(ctx context.Context, param *pb.Slice) (*pb.DeleteResponse, error) {
	var p pb.DeleteResponse
	it := index.Delete(param.Key)
	if it == nil {
		p.Success = false
		p.Value = []byte{}
	}
	p.Success = true
	p.Value = it
	return &p, nil
}

func main() {
	flag.Parse()
	if err := logger.SetupLogWithConf(*logConf); err != nil {
		panic(err)
	}
	defer logger.Close()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDDJkvServer(s, &server{})
	go until.Memery()
	table := index.NewTable("./data", "test", "test")
	f := table.CreateNewTable()
	tr = index.BuildBTreeFromPage(table.GetIndexPath(), f)
	s.Serve(lis)
}
