package main

import (
	"fmt"
	"reflect"
)

//通过反射,获取变量的类型

func main() {
	/*
		//a := 10
		//
		//b := 23.4
		//c := true
		//reflectFn(a)
		//reflectFn(b)
		//reflectFn(c)

		//var v  = Person{
		//	Name: "嘿嘿",
		//	Age: 123,
		//}
		//reflectFn(v)

		//var i *int
		//reflectFn(i)
	*/
	/*	//getValue(21)
		//getValue("wsn")
		//getValue(true)*/
	/*	var a int64 =  100
		setValue(&a)

		fmt.Println(a)*/

	var p1 = Person{
		Name: "嘿嘿",
		Age:  23,
	}
	getValueFromStruct(&p1)

}

//获取结构体里面的参数
func getValueFromStruct(s interface{}) {
	/*//获取变量类型
	ty := reflect.TypeOf(s)
	//获取变量值
	//val := reflect.ValueOf(s)

	//判断是否是一个结构体
	if ty.Kind() != reflect.Struct && ty.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数 不是一个结构体")
		return
	}

	val := reflect.ValueOf(s)
	var num int
	if ty.Kind() == reflect.Struct {
		num = ty.NumField()
		for i := 0; i < num; i++ {
			fmt.Printf("类型变量的名称:%v,", ty.Field(i).Name)
			fmt.Printf("类型变量的类型:%v,", ty.Field(i).Type)
			fmt.Printf("类型变量的json值:%v,", ty.Field(i).Tag.Get("json"))
			fmt.Printf("变量的值:%v\n", val.Field(i))
		}
	} else {
		num = ty.Elem().NumField()
		for i := 0; i < num; i++ {
			fmt.Printf("类型变量的名称:%v,", ty.Elem().Field(i).Name)
			fmt.Printf("类型变量的类型:%v,", ty.Elem().Field(i).Type)
			fmt.Printf("类型变量的json值:%v,", ty.Elem().Field(i).Tag.Get("json"))
			fmt.Printf("变量的值:%v\n", val.Elem().Field(i))
		}
	}*/

	//t:= reflect.TypeOf(s)
	//vv := t.Method(0)
	//fmt.Println(vv)

	//val := reflect.ValueOf(s)
	//vv:=val.MethodByName("XXX")
	//ss:=vv.Call(nil)
	//fmt.Println(ss)

	//方法的数量
	fmt.Println(reflect.TypeOf(s).NumMethod())
	fmt.Println(reflect.ValueOf(s).NumMethod())

}

//通过反射,设置变量的值 (指针)
func setValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println(v.Elem().Kind())
	v.Elem().SetInt(123)
}

//通过反射获取变量的原始值
func getValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind()
	switch kind {
	case reflect.Int:
		fmt.Println(v.Int() + 9)
		break
	case reflect.String:
		fmt.Println(v.String() + "hello")
		break
	case reflect.Bool:
		fmt.Println(!v.Bool())
		break
	default:
		fmt.Println("类型需要补全")
	}
}

func reflectFn(x interface{}) {
	//v :=reflect.ValueOf(x)
	v := reflect.TypeOf(x)
	fmt.Println(v)
	fmt.Println(v.Name())
	fmt.Println(v.Kind())
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Hobby string `json:"hobby"`
}

func (p Person) XXX() {
	fmt.Println("hello XXX")
}

func (p Person) GetMethod() {
	fmt.Println("hello getMethod")
}

func (p *Person) SetMethod() {
	fmt.Println("hello setMethod")
}
