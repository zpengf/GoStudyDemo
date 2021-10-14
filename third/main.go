package main

import "fmt"

func main()  {

	//指针 变量
	a:=1
	fmt.Println(a)
	fmt.Println(&a)

	//变量地址不变
	a = 2

	fmt.Println(&a)
	fmt.Println(a)

	//声明b 在 内存中开辟了一块空间
	//a只是给了b值 而不是地址
	b:=a
	fmt.Println(b)
	fmt.Println(&b)

	//指针变量
	//指针 指向内存空间
	//&就是指针

	c:=&a
	fmt.Printf("%T",c)

	//指针类型写法
	var d *int //没有任何指向  是个空指针
	//声明指针不会开辟地址 只是指向内存地址的某个地方
	//声明变量会开辟地址
	//指针变量接收的是一个变量的地址 否则没有办法告诉程序直接开辟空间

	e:=123

	d = &e


	*&e = 456 //就是*d
	//带&就是地址  带上*号就是取出地址中的值

	//可以*指针 可以获取地址中的值
	fmt.Println(*d)

	//应用指针 实现多个地方操作同一个地方的值  方法参数应用指针较多


	//--------------------new 函数

	//每次创建一个指针 必须额外创建一个变量 比较麻烦


	a1:=new(int) //这样直接就有了地址
	fmt.Println(a1)

	var a2 *int //这个就没有地址 是个空指针
	fmt.Println(a2)

	*a1 = 21

	*a2 = 21 //这个就会报错 空指针



}
