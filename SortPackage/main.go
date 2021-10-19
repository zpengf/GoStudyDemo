package main

import "sort"

func main() {


	//sort是个包 调用哪个排序 是程序内部决定

	num:=[]int{1,523,2,451,24,45,234}

	sort.Ints(num)//升序

	//升序
	sort.Sort(sort.IntSlice(num))


	//告诉sort 逆排序
	sort.Reverse(sort.IntSlice(num))

	sort.Sort(sort.Reverse(sort.IntSlice(num)))

	//应该是个升序排序 返回切片中的索引 如果不包含 则返回应该插入切片的哪个位置
	sort.SearchInts(num,1)



}
