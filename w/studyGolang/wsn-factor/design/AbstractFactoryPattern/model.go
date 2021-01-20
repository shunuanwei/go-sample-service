package AbstractFactoryPattern

// 原型复制抽象接口
type Prototype interface {
	clone() Prototype
}

type Message struct {
	Header *Header
	Body   *Body
}
type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}
type Body struct {
	Items []string
}

func (m *Message) clone() Prototype {
	msg := *m
	return &msg
}
