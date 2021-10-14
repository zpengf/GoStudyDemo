package main

import "fmt"

func main()  {
	//goto loop
	//goto 跳转到loop  标签名随意

	i:=6
	if i == 6 {
		goto Lable
	}
	fmt.Println("跳过")
	Lable:
		fmt.Println("程序结束")



}
