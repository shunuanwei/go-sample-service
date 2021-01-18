package main

import (
	"fmt"
	"net"
)

func main() {

	// http 模拟 get 请求
	/*	QQ, err := http.Get("https://www.baidu.com/")
		if err!= nil {

			fmt.Println(err.Error())
		}
		defer QQ.Body.Close()

		all, _ := ioutil.ReadAll(QQ.Body)

		fmt.Println(string(all))*/

	// http 模拟 post请求
	/*	post, err := http.Post("", "application/json;charset=utf-8", bytes.NewBuffer([]byte("")))*/

	//http.Client和http.NewRequest来模拟请求
	/*	v := url.Values{}
		v.Set("username", "xxxx")
		v.Set("password", "xxxx")
		//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
		body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
		client := &http.Client{}//客户端,被Get,Head以及Post使用
		reqest, err := http.NewRequest("POST", "http://xxx.com/logindo", body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}
		//给一个key设定为响应的value.
		reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded;param=value") //必须设定该参数,POST参数才能正常提交

		resp, err := client.Do(reqest)//发送请求
		defer resp.Body.Close()//一定要关闭resp.Body
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
		}

		fmt.Println(string(content))*/

	//创建简单 服务器监听

	addr, err := net.ResolveTCPAddr("tcp", "9078")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var request []byte
	for true {
		tcp, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		_, _ = tcp.Read(request)
		fmt.Println(string(request))
		tcp.Close()
	}


	net.Dial(



}
