package interfacePkg

import "fmt"

//接口中的方法,需要使用 结构体 或者 自定义类型 实现

type Phone struct {
	Name  string
	Num   string
	Price int
}

//如果是 值接受者   值和指针都可传入
//如果是 指针接受者    只能使用 指针去使用
func (p Phone) start() {
	fmt.Println(p.Name + "启动------")
}

func (p Phone) close() {
	fmt.Println(p.Name + "关机------")
}

type Computer struct {
	Name  string
	price int
}

func (c Computer) Work(u Usber) {
	u.start()
	u.close()
}
