package _defer

import "fmt"

//完成延迟功能  当函数执行完 执行defer功能


//常用就是  关闭连接 数据库连接 文件

func main() {
	fmt.Println("open correct")
	defer func() {
		fmt.Println("连接关闭")
	}()

	fmt.Println("执行")

}


//多个defer
// 最后的defer先执行  栈结构

//defer return 存在时  先给return的值赋值
// 然后 defer

