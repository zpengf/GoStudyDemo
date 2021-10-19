package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {

	//go 中 有双向链表 循环链表

	//前连接 内容 后连接

	//头元素 前连接为nil  后连接考虑后边是否有元素
	//尾元素 后连接为nil

	//充分利用内存空间 实现内存灵活管理

	//缺点 空间开销比较大  遍历时跳跃查找 大量数据效率比较低


	//实例话 list
	list1:=list.New()
	fmt.Println(list1)

	//前边加一个
	list1.PushFront("a")

	//最后添加一个
	list1.PushBack("b")

	//在最后元素后边插入c
	list1.InsertAfter("c",list1.Back())

	list1.InsertBefore("d",list1.Front())

	//遍历
	for e:=list1.Front();e!=nil;e = e.Next(){
		fmt.Println(e.Value)
	}

	//移动元素的顺序

	//把最后一个移动到第一个前边
	list1.MoveBefore(list1.Back(),list1.Front())

	//-----------------------------双向循环链表

	//没有头元素 没有尾元素
	//go 会把第一个创建的元素 当作头元素

	r:=ring.New(5)
	//r代表第一个元素
	r.Value = 0
	r.Next().Value = 1

	//循环
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

	//向后移动几次 n<0 向前移
	//n>0 向后移
	r.Move(0)








}
