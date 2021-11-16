package xmlDir

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

//xml 数据存储 数据交互 配置文件
//跨平台性  数据工整
//第一行 xml头 版本 编码
//第二行 dtd 文件类型定义 语法检查
//最外层标签 根结点 每一层标签 元素节点
//标签有属性
//严格区分大小写
//标签正确嵌套 属性值有双引号


//读取xml
// 提供了对xml 序列化（把内容持久化读取到硬盘） 和反序列化（只读取）
//Unmarsha 直接转化成结构体

type Peoples struct {
	XMLName xml.Name `xml:peoples`
	Peos []People `xml:people`
}



type People struct {
	XMLName xml.Name `xml:"people"`
	Id int `xml:"id,attr"`
	Name string `xml:"name"`
	Address string `xml"address"`
}


func main() {
	peo := new(Peoples)
	b,_:=ioutil.ReadFile("demo.xml")
	xml.Unmarshal(b,peo)
	fmt.Println(peo)

	//xml文件生成  把结构体数据生成文件
	p := People{Id: 123,Name: "zhangsan",Address: "ceshi"}
	//b1,_:=xml.Marshal(p)
	b1,_:=xml.MarshalIndent(p,"","		")
	b1 = append([]byte(xml.Header),b...)
	ioutil.WriteFile("/users",b1,0777)
	fmt.Println("结束")








}


