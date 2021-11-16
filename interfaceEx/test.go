package interfaceEx

import "fmt"

//接口是一组行为规范的定义
//接口中只能有方法声明 方法只能有名称 参数 返回值
//每个接口 可以有多个方法声明 结构体要把接口的所有的方法都重写后 结构体属于接口类型
// interface

type People struct {
	name string
	age int
}

type Live interface {
	run(name string)
	eat()
}

func (peo *People) run(name string) {
	fmt.Println(peo.name + "正在跑步")
}
func (peo *People) eat() {
	fmt.Println(peo.name)
}

//多态
//同一件事 在不同时候结果不同

//接口 当作 方法的参数

//把  结构体类型 赋给接口 前提结构体必须全部重写了 接口的方法















