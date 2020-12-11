package interfacePkg

type Usber interface {
	start()
	close()
}

//空接口是一种数据类型 空接口可以标识任意类型
type All interface{}
