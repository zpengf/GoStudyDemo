package main

import (
	"fmt"
	"net/http"
)

func main() {

	//先写响应内容
	http.HandleFunc("/",welocme)

	//在开启占用端口监听
	err:=http.ListenAndServe("localhost:8081",nil)
	if err != nil {
		fmt.Println("监听出错：",err)
	}



}

func welocme(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type","text/html;charset=utf-8")
	fmt.Fprintf(w,"服务器返回给你第一个响应<b>您好</b>")

}
