package SingletonPattern

import (
	"fmt"
)

func main() {

	//放
	msgPool.Instance().AddMsg("hello")

	//取
	msg := msgPool.Instance().GetMsg()
	fmt.Println(msg)
}

//// 运行结果
//=== RUN   TestMessagePool
//--- PASS: TestMessagePool (0.00s)
//PASS
