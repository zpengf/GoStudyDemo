package main

import "fmt"

func main() {
	//map[key]value

	//map 值类型 声明没有内存地址
	var m map[string] int

	fmt.Println(m == nil)

	//map 并不是并发安全的
	//实例化
	m1:=make(map[string]string)

	//声明 map 直接赋初值
	m2:=map[string]string {"dd":"sss","ddd":"222"}

	//使用key 如果存在则修改value的值
	//key不存在 表示新增操作

	m2["ssss"] = "qqqqqq"

	//删除  key存在表示删除 删除不存在的不会报错也没有提示
	delete(m1,"ss")

	//循环遍历
	for key,value := range m2 {
		fmt.Println(key,value)
	}

}
