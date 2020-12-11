package main

import "fmt"

func main() {

	//var p1 = Phone{
	//	Name:"华为麦芒六",
	//}
	//var i1 Usber
	//
	//i1 = p1
	////i1.Start()
	//
	//var computer1 Computer
	//computer1.Work(i1)
	//computer1.Work(p1)

	//m :=make(map[string]interface{})
	//m["name"] = "嘿嘿嘿~~"
	//m["age"] = 13
	//fmt.Printf("%#v",m)

	judge("123")
	judge(123)
	judge(false)
}

func judge(x interface{}) {
	//x.(type)  只能在switch 中使用
	switch x.(type) {
	case string:
		fmt.Println("string 类型")
		break
	case bool:
		fmt.Println("bool 类型")
		break
	case int:
		fmt.Println("int 类型")
		break
	default:
		fmt.Println("没有找到")
	}
}
