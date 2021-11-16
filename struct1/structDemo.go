package struct1

import "fmt"

//结构体  java中类本质就是结构体
//结构体是个值类型

//结构体定义到哪里 就是哪里能用 定义到函数内部 就是个局部
//首字母 大小 挎包能访问
type Person struct {
	name string
	age int
}

type  Student struct {
	studentNo string
	Person
}
//结构体使用方法 函数属于包

//实现类型 继承  继承是实现代码复用的非常重要的手段

type Teacher struct {
	teaNo string
	class string
	Person
}

func (s *Student) SetPerson(name string,age int) {
	s.Person.age = age
	s.Person.name = name

}


//func(s *Student) SetName(name string){
//	s.name = name
//}
//
//func(s *Student) GetName() string{
//	return s.name
//}


//结构体指针
//通过new函数创建结构体指针

func test1()  {

	stu := new(Student)

	stu1 := stu

	fmt.Println(stu1 == stu)

	stu.name = "111"

	stu1.name = "222"
	stu1.age = 11





}




