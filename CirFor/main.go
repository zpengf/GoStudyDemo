package main

import "fmt"
/**
	for循环
 */
func main()  {
	//for 循环是go中唯一一个循环结构

	for i := 0;i < 8;i++{
		fmt.Println(i)
	}

	//遍历数组

	arr := [4]string {"","","",""}

	for i := 0;i < len(arr);i++{


	}
	//结合range
	//i 表示坐标 n表示内容
	for i,n := range arr{
		fmt.Println(i)
		fmt.Println(n)
	}

	//双重for循环
	arr2 := [5]int{1,5,32,5,2}

	for i:=0;i < len(arr2) - 1;i++{
		for j:= 0;j < len(arr2)-i-1;j++ {
			if arr2[j] > arr2[j+1]{
				arr2[j],arr2[j+1] = arr2[j+1],arr2[j]
			}
		}
	}

	fmt.Println(arr2)

	//continue break














}
