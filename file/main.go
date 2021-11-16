package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//file 既包括文件 也是文件夹

//file filemode  fileinfo

//相对路径 从 gopath 开始找

//绝对 从磁盘根目录开始找

func main() {
	//两种创建文件
	err:=os.Mkdir("/users/pengfeizhang/goTest",os.ModeDir)
	if err != nil{
		fmt.Println("文件夹创建失败",err)
		return
	}
	os.Chmod("goTest", 0777)

	fmt.Println("文件夹创建成功")

	//os.Create()


	//重命名
	//os.Rename()


	//文件信息
	//os.Open()

	//删除内容
	//os.Remove()

	//删除目录
	//os.RemoveAll()



	//reader 输入流  文件读取到应用文件 使用流
	//一切都相对于程序来说

	r:=strings.NewReader("ssss")

	//生成一个切片
	b:=make([]byte,r.Size())

	//返回值是个长度 和 error
	n,err:=r.Read(b)

	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("长度",n)

	//读取文件流
	f,err:=os.Open("/users/pengfeizhang/ss.txt")

	finfo,err:=f.Stat()

	//新建一个切片去接收文件里的内容
	b1:=make([]byte,finfo.Size())

	f.Read(b1)


	fmt.Println(string(b1))



	//writer  输出流
	//使用输出流的 不能用os.open 这是只读
	//os.OpenFile



	//ioutil包
	//readAll 读取内容直接返回切片
	f1,_:=os.Open("/users/go.txt")
	defer f1.Close()
	b12,_:=ioutil.ReadAll(f1)
	fmt.Println(string(b12))







}
