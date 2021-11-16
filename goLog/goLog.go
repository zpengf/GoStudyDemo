package goLog

import (
	"log"
	"os"
)

//日志信息打印 log包

//print() 输出日志信息

func main() {

	log.Println("打印日志信息")

	log.Panic() //这个会让程序停止

	log.Fatal("打印feat信息 会立即终止程序")

	//｜ 表示多种情况
	f,_:=os.OpenFile("/users/go.log",os.O_APPEND|os.O_CREATE,07777)

	//f是流  第二个参数是前缀
	logger:=log.New(f,"[info]",log.Ltime)
	logger.Println("打印日志信息")



}


