package main

import (
	"fmt"
	"math"
	"time"
)

func main()  {
	//if语句
	s:=65
	if s >= 69 {
		fmt.Println("及格")
	}


	//直接在判断处 写函数
	if a:=65; a >= 45 {
		fmt.Println(a)
	}


	//-------switch
	num:=12
	switch num {
	case 1:
		fmt.Println("ss")
	case 2:
		fmt.Println("ss")
	default:
		fmt.Println("default")

	}

	//--
	num=12
	switch {
	case num >= 2:
		fmt.Println("aa")
	default:
		fmt.Println("default")
	}

	//----
	num=12
	switch num {
	case 1,2,3,4,5:
		fmt.Println("aa")
	default:
		fmt.Println("default")
	}

	//fallthrough 穿透
	num=12
	switch num {
	case 1,2,3,4,5:
		fmt.Println("aa")
		fallthrough//执行下个case 只往下执行一个case
	case 22,32,13,44,54:
		fmt.Println("aa")
	default:
		fmt.Println("default")
	}

	//break 直接跳出循环


	//------------------------时间类型
	//默认时间是无意义的
	var t time.Time
	fmt.Println(t)

	//获取当前时间
	t = time.Now()
	//一般用纳秒
	t = time.Unix(0,t.UnixNano())

	//指定时间
	time.Date(2022,2,14,12,12,3,4,time.Local)

	t.Year()
	t.Month()
	t.Day()
	t.Date() //输出日期

	t.Hour()
	t.Minute()
	t.Second()
	t.Nanosecond()

	t.Clock()//直接输出时间

	t.Unix()//秒差 距离1970年1月1日的时间差

	t.UnixNano()//纳秒 差

	//时间与string 相互转化

	t = time.Now()
	t.Format("2006-01-02 15:04:06")

	//常用的数学函数
	var m1,m2 float64 = 12.3,13.5

	//向下取整
	math.Floor(m1)

	//向上取整
	math.Ceil(m2)

	//取绝对值
	math.Abs(12.4)


























}
