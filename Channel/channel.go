package Channel

import "fmt"

//线程通信是难点 go中提供了语言级别的通信 channel 管道 信道 通道

//channel是进程内的通信方式 每个channel 只能传递一个类型的值
//声明channel时指定

//两个作用
//通信  同步

//声明channel
//channel 是安全的  多个协程同一时间操作只能又一个
func main() {
	var c1 chan int //可读可写
	var c2 chan <- int //只写
	var c3 <- chan int //只读
	c4:=make(chan int)//无缓存channel
	c5:=make(chan int,0)//无缓存
	c6:=make(chan int,100)//有缓存

	ch:=make(chan int)

	go func() {
		fmt.Println("start")


		fmt.Println("end")


		//如果没人取 那就会一直停在这里
		ch <- 998

	}()

	a:=<-ch
	fmt.Println(a)
	fmt.Println("程序结束")

	//无缓存添加内容 或者添加的内容个数大雨channel缓存个数就会产生死锁

	//主协程阻塞 会产生死锁


}



