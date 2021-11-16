package duanyan

import "fmt"

//只要实现了 接口的全部方法 就认为这个类型属于这个接口
//如果一个接口中没有任何方法 则所有类型都实现了这些接口

//interface{} 作为方法参数 就可以接收任意类型

//断言 判断参数是什么类型

func demo(i interface{})  {

	result:=i.(int)

	fmt.Println(result)
}

func main() {
	var i interface{} = 123;

	fmt.Println(i.(int))
}