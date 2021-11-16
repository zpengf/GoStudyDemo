package _select

import (
	"fmt"
)

//跟 switch 很像 case条件只能是io  channel的数据存取就是io操作

func main() {
	ch1 := make(chan int,1)
	ch2 := make(chan string)

	ch1 <- 1
	ch2 <- "ss"

	select {
	case a1:= <-ch1:
		fmt.Println(a1)
	case a2:= <-ch2:
		fmt.Println(a2)


	}




}



