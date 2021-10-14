package main

import "fmt"

func main()  {

	//数组 固定长度 相同类型的序列

	var arr [3]int = [3]int{1,2,3}

	//初始值 为0
	fmt.Println(arr)

	fmt.Println(arr[0])

	//arr1:=[5]int{}

	//数组长度 看内存空间

	arr2 :=[...]int{1}

	fmt.Println(len(arr2))

	//数组 在go中是值类型
	//把数组变量给另一个数组变量 会重新开辟一块空间 值一样 地址不一样
	//两个数组是完全独立的数组

	//--------------多维数组
	//一个数组变量中每个元素又是一个数组变量


	//var arr4 [3][3]int
	//第一个3表示行 第二个表示列

	arr4 := [3][3]int{
		{1,2,3},
		{3,3,3},
		{3,3,3},
	}

	fmt.Println(arr4)


























}
