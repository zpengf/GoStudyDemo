package recover

import "fmt"

//恢复程序的正常运行 接收panic的信息

func main() {
	defer func() {
		recover() //就不会输出panic中的消息
		if error:=recover();error !=nil{
			fmt.Println(11)
		}

	}()

	fmt.Println(111111)

	panic("程序终止")

	fmt.Println(222222222)
}