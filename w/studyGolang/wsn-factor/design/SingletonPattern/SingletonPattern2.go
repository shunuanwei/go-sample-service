package SingletonPattern

import "sync"

//单例模式 - 懒汉模式

var once = &sync.Once{}

// 消息池单例，在首次调用时初始化
var msgPool2 *messagePool

// 全局唯一获取消息池pool到方法
func Instance() *messagePool {
	// 在匿名函数中实现初始化逻辑，Go语言保证只会调用一次
	new(sync.Once).Do(func() {
		msgPool2 = &messagePool{
			// 如果消息池里没有消息，则新建一个Count值为0的Message实例
			pool: &sync.Pool{New: func() interface{} { return "&Message{Count: 0}" }},
		}
	})
	return msgPool2
}
