package main

import "fmt"

func main()  {
	demo1(11)
	add(1,2)

	add2(1,2)

	a,b:=add2(1,2)

	c,_:=add2(1,2)
	fmt.Println(a,b)
	fmt.Println(c)
}

func demo1(a int)  {
	fmt.Println(a)
}
func add(a,b int)  {
	fmt.Println(a+b)
}
func add1(a,b int) int {
	return a+b
}
//多返回值函数
func add2(a,b int) (int,int) {

	//接的时候 必须用两个接
	//如果实在不想要 可以用_占位置

	return a,b

}







