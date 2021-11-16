package WaitGroup

import (
	"fmt"
	"sync"
)

//go sync 提供了基本同步基元

var (
	num = 10
	wg1 sync.WaitGroup
	//声明互斥锁
	m sync.Mutex
)


func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i:=0;i < 5;i++{
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()



}
func dem1()  {
	//互斥锁 go run -race 终端输入这个可以查看数据有没有竞争
	//读写锁  map就不是线程安全的 go 去访问map就会发生错误


	//多个协程 操作一个变量出现冲突问题 竞争
	m.Lock()
	for i := 0;i<100;i++{
		num = num - 1
	}
	m.Unlock()
	wg1.Done()
}

