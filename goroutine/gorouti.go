package goroutine

import (
	"fmt"
	"time"
)

//go 从语言层面就支持并发

// goroutine 就是协程 类似于其他语言线程 但是比线程更轻
//main 主函数就是个协程
//其他协程 都是听主函数 主函数完事 整个程序结束

//起到分流作用  可以让程序同时具备同时做多件事进行

//并行 不同代码片段同时在不同的物理处理器上支持

//并发 同时管理多个事情 物理处理器可以运行某个内容一半后处理其他事情

//并发是以较少的资源去做更多事情

//主流并发模型
//多线程
//基于回调的异步io
//协程 不需要抢占式调用

func demo(cout int)  {
	for i:=0;i<100;i++{
		fmt.Println(cout,i)
	}
}

func main() {
	for i:=0;i<5;i++ {
		//开启了5个协程
		go demo(i)
	}
	time.Sleep(3e9)
}



