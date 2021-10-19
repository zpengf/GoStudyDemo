package main

import "fmt"

func main() {
	//切片 长度可变的相同类型元素序列  有点像 规定类型集合？？
	//声明切片
	var slice []string

	//声明类型
	var array [2]string

	//切片相当于一个指针 只声明时没有 开辟内存空间 不能直接操作 是nil
	//输出内存地址为0x0 没有地址

	slice1 :=[]string {"1","2"}

	slice = []string {"11","23"}

	fmt.Println(slice1)

	//切片是引用类型 两个相等 内存地址是相等的 用=时 把地址给新的

	// 切片 make函数 创建slice map channel interface
	//使用make函数定义 意味着切片已经申请了内存空间
	//make(类型,初始长度,初始容量) 长度是个数 容量是申请多少空间 容量可省略
	s:=make([]int,0)

	fmt.Printf("%p\n",s)

	//len查看长度 cap查看容量
	fmt.Println(len(s))
	fmt.Println(cap(s))
	//append() 直接用的顶级函数 一次增加多个值 但是必须切片接收
	s = append(s,2)
	s = append(s,3)
	s = append(s,4)

	//容量在超出上一次容量值 会翻倍 如果扩容的大于上一次容量 则容量和长度相等
	s1:=[]int{1,2,34}
	s = append(s,s1...)

	//通过数组产生切片 定义数组后 取出一个片段 取出的就是切片
	arr := [4]int{1,2,34,5}
	s2:=arr[0:2]
	fmt.Println(s2)
	//s2 就是一个切片
	//切片就是个指针 指向内存地址 修改切片内容  数组也会变
	s2[0] = 23232
	fmt.Println(arr)
	//切片增加内容 就是在切片后边增加内容
	//会改变数组原有内容 改了切片后边的地址 切片后边的地址数组已经用了
	//如果超出了数组长度 会重新给切片申请一块空间 因为原来在数组后边的空间没有被申请

	//通过数组产生切片 需要注意的就是 切片就是个指针

	//删除的实现 通过切片取其中一小块形成新的切片
	//删除索引为2的
	delindex:=2
	//取出前三个 将后边的添加到新的切片
	newS := s[0:delindex]
	newS = append(newS,s[delindex+1:]...)


	//copy函数 原切片内容copy到目标切片 按角标 原角标0 到新的角标为0

	scopy1:=[]int{1,2}
	scopy2:=[]int{3,4,3,4}

	copy(scopy2,scopy1)
	//结果 scopy2 = [1,2,3,4]
	//s2 copy到 s1  s1长度不会变 只会角标1 2的值会变


	//使用copy删除元素 删除后原切片不会变
	newSilce1 :=make([]int,1)
	copy(newSilce1,scopy1[0:1])
	newSilce1 = append(newSilce1,s[1+1:]...)




































}
