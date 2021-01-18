package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

func main() {

	saramaProduce()
}

// 生产者
func saramaProduce() {

	config := sarama.NewConfig()

	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于v0_10_0版本,消息中timestrap没有作用,需要消费和生产同时配置
	//注意:版本设置不对的话,kafka会返回特殊的错误,并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	config.Net.SASL.Enable = true
	config.Net.SASL.User = "admin"
	config.Net.SASL.Password = "123456"

	fmt.Println("开始配置生产者")

	producer, err := sarama.NewAsyncProducer([]string{"121.40.115.212:19192"}, config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer producer.AsyncClose()

	//循环判断是哪个通道发送过来的数据
	fmt.Println("开启循环: 监听数据")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				fmt.Println("err : " + fail.Error())
			}
		}
	}(producer)

	var value string
	for i := 0; ; i++ {
		time.Sleep(5000 * time.Millisecond)
		time11 := time.Now()
		value = "this is a message 0606 ==>> " + time11.Format("15:04:05")

		// 发送的消息,主题。
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: "0606_test",
		}

		//将字符串转化为字节数组
		msg.Value = sarama.ByteEncoder(value)
		//fmt.Println(value)

		//使用通道发送
		producer.Input() <- msg
	}

}
