package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	pb "github.com/sdgmf/go-sample-service/w/studyGolang/wsn-grpc/grpc_net/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"net"
	"strconv"
	"time"
)

/**
服务器端,需要类似于重写protobuf 里面的方法,方法名称 必须一样,所有方法 都必须存在
*/

//定义一个空的结构体,为了与pb.go中的service实体对应
type SService struct{}

func (ss *SService) SeaHello(ctx context.Context, animal *pb.Animal) (*pb.AnimalResponse, error) {
	fmt.Println("StudentAdd================简单模式================")
	ret := &pb.AnimalResponse{Code: 200, Data: animal.String()}
	return ret, nil
}

func (ss *SService) GetStudent(stream pb.AnimalService_GetStudentServer) error {
	fmt.Println("服务器端,双向模式 === >>")
	for {
		//接收
		in, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("read done")
			return nil
		}
		if err != nil {
			grpclog.Fatal("some error in recv", err)
		}
		fmt.Println(in, strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
		//发送
		rp := &pb.AnimalResponse{Code: 1000, Data: "n xin"}
		_ = stream.Send(rp)
	}
	return nil
}

func main() {

	fmt.Println("Service : ===>>")
	listen, err := net.Listen("tcp", "127.0.0.1:9077")
	if err != nil {
		log.Error("服务器端创建监听失败")
		return
	}
	server := grpc.NewServer()
	pb.RegisterAnimalServiceServer(server, &SService{})
	_ = server.Serve(listen)
}
