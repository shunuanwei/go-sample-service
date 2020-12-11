package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func mai3n() {
	//filePath1 := "E:\\GO\\hello.txt"
	//filePath2 := "E:\\GO\\hello2.txt"
	//复制文件

	//1 -- ioutil
	//	file, err := ioutil.ReadFile(filePath1)
	//	if err != nil {
	//		fmt.Println("读取文件失败")
	//		return
	//	}
	//	err = ioutil.WriteFile(filePath2, file, 0666)
	//	if err != nil {
	//		fmt.Println("文件写入失败")
	//	}

	/*	//2 -- 文件流的方式
		read, err := os.Open(filePath1)
		write, err1 := os.OpenFile(filePath2, os.O_APPEND, 0666)
		defer read.Close()
		defer write.Close()
		if err != nil || err1 != nil {
			fmt.Println(err)
			return
		}
		tempSlice := make([]byte, 128)
		for true {
			n1, e1 := read.Read(tempSlice)
			if e1 == io.EOF {
				break
			}
			if e1 != nil {
				fmt.Println("读取文件失败")
				return
			}
			_, e2 := write.Write(tempSlice[:n1])
			if e2 != nil {
				fmt.Println("写入失败")
				return
			}
		}*/

}

func main() {
	write, err1 := os.OpenFile("E:\\look\\不死者之王.txt", os.O_CREATE|os.O_APPEND, 0666)
	if err1 != nil {
		panic("打开创建文件失败")
	}
	defer write.Close()

	tempSlice := make([]byte, 128)
	for i := 1; i <= 14; i++ {
		filePath1 := "E:\\look\\" + strconv.Itoa(i) + ".txt"
		read, err := os.Open(filePath1)
		if err != nil {
			panic("打开文件失败" + filePath1)
		}
		for true {
			n1, e1 := read.Read(tempSlice)
			if e1 == io.EOF {
				break
			}
			if e1 != nil {
				fmt.Println("读取文件失败")
				return
			}
			//fmt.Println(n1)
			//fmt.Println(string(tempSlice))
			_, e2 := write.Write(tempSlice[:n1])
			if e2 != nil {
				fmt.Println("写入失败")
				return
			}
		}
		read.Close()
	}

}
