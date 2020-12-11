package main

import (
	"io/ioutil"
)

func main() {

	filePath := "E:\\GO\\hello.txt"
	//读取文件
	//1 --  file.read
	//	file, err := os.Open(filePath)
	//	defer file.Close()
	//	if err != nil {
	//		fmt.Println("文件打开失败")
	//	}
	//	var ret []byte
	//	var bb = make([]byte,1024)
	//	for {
	//		num,err:=file.Read(bb)
	//		if err == io.EOF {
	//			fmt.Println("读取完毕")
	//			break
	//		}
	//		if err != nil {
	//			fmt.Println("读取文件失败")
	//			return
	//		}
	//		ret = append(ret, bb[:num]...)
	//	}
	//
	//	fmt.Println(string(ret))

	//2 --- bufio
	//	file,err:=os.Open(filePath)
	//	defer file.Close()
	//	if err != nil {
	//		fmt.Println("打开文件失败")
	//		return
	//	}
	//	var ssss string
	//	//读取文件
	//	reader := bufio.NewReader(file)
	//	for true {
	//		//按行读取
	//		str, err:= reader.ReadString('\n')
	//		if err == io.EOF {
	//			fmt.Println("读取完毕")
	//			ssss = ssss + str
	//			break
	//		}
	//		if err != nil {
	//			fmt.Println("读取文件失败")
	//			return
	//		}
	//		ssss = ssss + str
	//	}
	//	fmt.Println(ssss)

	//3 -- ioutil
	//	bb, err := ioutil.ReadFile(filePath)
	//	if err != nil{
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(string(bb))

	//写入文件
	//1 -- file.w 使用 file
	//	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	//	defer file.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	//写入文件
	//	writeString, err := file.WriteString("hello world XXX \r\n")
	//	if err != nil{
	//		fmt.Println("文件写入失败")
	//		return
	//	}
	//	fmt.Println(writeString)

	//2  bufio
	//	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	//	defer file.Close()
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	writer := bufio.NewWriter(file)
	//	writeString, err := writer.WriteString("今天是周一")
	//	_ = writer.Flush()
	//	if err != nil {
	//		fmt.Println("写入失败")
	//		return
	//	}
	//	fmt.Println(writeString)

	//3 -- ioutil   没有append
	_ = ioutil.WriteFile(filePath, []byte("明天是周二"), 0666)

}
