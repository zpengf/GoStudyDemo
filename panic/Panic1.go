package panic

import (
	"fmt"
	"os"
)

//当执行到panic 就会终止当前代码 打印错误信息

func main() {

	fmt.Println(111111)
	panic("程序终止")

	fmt.Println(22222222)

	//并不是立即停止程序运行  比如会执行defer中

	os.Exit(0)//程序立即终止

}
