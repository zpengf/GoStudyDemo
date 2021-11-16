package ProceedingSleep

import (
	"fmt"
	"time"
)

func main() {

	//main 函数是主线程 协程  整个程序从上到下执行

	//参数 是纳秒 e表示10的9次方 1e9就是1s

	//阻塞程序1秒
	time.Sleep(1e9)

	//延迟执行
	time.AfterFunc(3e9,func(){
		fmt.Println("程序延迟执行")
	})

	time.Sleep(5e9)

	fmt.Println("程序结束")


}
