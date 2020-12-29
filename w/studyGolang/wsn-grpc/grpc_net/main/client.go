package main

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	pb "github.com/sdgmf/go-sample-service/w/studyGolang/wsn-grpc/grpc_net/proto"
	"google.golang.org/grpc"
	"io"
)

/**
客户端,只需要写此客户端需要用到的方法就好,最好也是全部方法都存在
*/

func SeaHello(client pb.AnimalServiceClient, animal *pb.Animal) {
	fmt.Println("client == > SeaHello================简单模式================")
	//简单模式下直接调用方法就可以传递参数和获得返回值
	hello, err := client.SeaHello(context.Background(), animal)
	if err != nil {
		log.Error("error")
		return
	}
	fmt.Println("sea Hello :", hello.Code, hello.Data)
}

func GetStudent(client pb.AnimalServiceClient) {
	fmt.Println("client : GetHelloTest_Method================客户端模式================")
	//此处相当于创建了一个GetHelloTest的通道。通过通道发送和接收数据。
	re, err := client.GetStudent(context.Background())
	if err != nil {
		fmt.Println("some error occur in gethellotest", err)
	}
	//初始化要发送给服务器端的数据，并发送（此处连续发送10次）
	for i := 0; i < 10; i++ {
		sq := &pb.Animal{Age: 1}
		re.Send(sq) //客户端要先发送数据再接收
	}
	//使用for循环接收数据
	for {
		//先关闭send然后再接收数据。该方法内部调用了CloseSend和RecvMsg方法。但是后者协程不安全
		r, err2 := re.Recv()
		if err2 == io.EOF {
			fmt.Println("recv done")
			return
		}
		if err2 != nil {
			fmt.Println("some error occur in gethellotest recv")
		}
		fmt.Println(r)
	}

}

func main() {
	fmt.Println("client ===>>")
	dial, err := grpc.Dial("127.0.0.1:9077", grpc.WithInsecure())
	if err != nil {
		log.Error("client dial error")
		return
	}
	defer dial.Close()
	client := pb.NewAnimalServiceClient(dial)
	s := pb.Animal{Name: "huang", Age: 3}

	SeaHello(client, &s)

	GetStudent(client)
}
