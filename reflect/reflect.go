package reflect

import (
	"fmt"
	"reflect"
)

//运行时反射 动态操作结构体
//动态的获取

type People struct {
	name string `xml:aaa`
	Address string
}

func main() {

	a:="name"
	//获取类型
	reflect.TypeOf(a)
	//获取内容
	reflect.ValueOf(a)

	//获取结构体属性的值
	peo:=People{"zhangsan","wojia"}

	v:=reflect.ValueOf(peo)

	//属性个数
	v.NumField()
	//获取属性内容 获取第二个
	v.FieldByIndex([]int{1})
	//根据名字
	v.FieldByName("address")

	//设置属性的值 设置值 必须指针
	peo1:=new(People)
	// 地址的值必须经过elem才能操作
	v1:=reflect.ValueOf(peo1).Elem()

	//查看是否可一设置
	//当我们需要修改结构体 属性名时  结构体属性名必要大写
	v1.FieldByName("address").CanSet()

	v1.FieldByName("address").SetString("ceshi")


	//标记
	//获取标记

	t:=reflect.TypeOf(People{})
	name1,_:= t.FieldByName("name")
	tat := name1.Tag
	fmt.Println(tat)




}

