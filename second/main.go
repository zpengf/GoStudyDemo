package main

import (
	"fmt"
	"strconv"
)

func main()  {

//	var name string
//
//	name = "ss"
//
//	fmt.Println(name)
//
//
//	var a string = "1"
//
//	fmt.Println(a)
//
//	var s = "111"
//	fmt.Println(s)
//
//	b:=a
//	fmt.Println(b)
//
//
////一次声明多个变量
//	var e,er,we int
//	e,er,we = 1,2,3
//	fmt.Println(e,er,we)
//
//	var (
//		z = 1
//		v = 2
//		x = false
//	)
//	fmt.Println(z,v,x)


	//整型
	//数学运算的类型 分为 整型 浮点型

	//int8 和 int16不能进行运算
	//相同类型才能运算
	//有符号整型 无符号整型
	//二进制 最高位 0为正数 1为负数
	//
	/**
	  	int8 取值范围 -128 ---  127
		int16 -32768 32767
		int 受限于计算机操作类型 int为多少位

		uint8 0 255
	 	uint16 0 65535


		byte 与uint8类似 一个字节占8个二进制

		ruin
	 */

	//类型转换 运算必须保证同一类型
	//go中 int8 不能自动转换成int16
	var a int8 = 8
	var b int16 = 16
	var c int32 = 32
	var d int64 = 64
	var e int = 10
	//go中默认是int类型
	d:=11

	fmt.Println(a + b)

	//类型转化同一语法 类型（值）
	fmt.Println(int16(a) + b)

	//en 表示10的n次方


	//-===========================字符型
	//go语言不支持字符型类型 go会转换为对应编码表中int32的值

	//字节中前32为控制吗 后在控制码基础上添加空格 标点符号 大小字母到127 形成ACSII吗


	//float64 误差几率小
	//float32 误差几率大

	//字符串
	var str string = "smallling"
	str1:="ss\ns"
	//\n代表换行

	//str2:="111"
	fmt.Println(str+str1)

	//string converse
	//字符串转数字 10表示当前是十进制数 64表示 int64
	strconv.ParseInt(str,10,64)

	//默认 10进制 int64
	strconv.Atoi(str)

	//int转成str 10表示 10进制 返回类型是个字符串
	strconv.FormatInt(12,10)

	strconv.Itoa(12)

	//字符串截取
	//len 获取字符串字节长度 英文占1个字节长度 中文占3个字节长度
	str3:="adf长"
	a22:=len(str3)


	//常量 常量在编译期就确定 不可以改变
	//const 驼峰写法
	const con1 = 123
	const con3 = "21"






































}
