package expectiong

import "errors"

//go 中错误作为方法 函数的返回值

//自定义错误
func demo(i,k int) (r int,e error) {

	if k == 0{
		e = errors.New("no 0")
		return
	}
	r= i/ k
	return
}




