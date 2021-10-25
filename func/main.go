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

	//匿名函数 //没有名的函数
	re := func (a int) (b int) {
		fmt.Println(a)
		b = a + 1
		return b
	}(1)//括号表示调用
	fmt.Println(re)

	//函数也是一种类型 是引用类型 赋值传递的是内存地址的值
	var func1 func()
	var func2 func(int)
	var c func(int) int

	fmt.Println(func1)

	//函数作为参数
	mydo(func(name string) {
		fmt.Println(name)

	})

	res:=func222(1)
	resss := res(2)
	fmt.Println(resss)

}
func mydo(arg func(name string))  {
	fmt.Println("aaa")
	arg("aagaga")
}

//函数作为返回值
func func222(a int) func(b int) int {

	a = a + 1

	return func(b int) int {
		return a + b
	}
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


//可变参数函数 调用函数时  参数的个数可是任意个
//其实传的就是个 就是个切片类型
func changeFunc(name string,b... int)  {
	for i,n:=range b {
		fmt.Println(i,n)
	}
}








