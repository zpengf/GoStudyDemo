package main

import (
	"database/sql"
	"fmt"
	//导入驱动
	_"github.com/go-sql-driver/mysql"
)

func main() {

	//root 用户名 后边是密码
	db,err:=sql.Open("mysql","root:root@tcp(localhost:3306)/db_file")

	//打开数据库连接
	db.Ping()

	defer func() {
		if db != nil{
			db.Close()
		}
	}()

	if err!=nil{
		fmt.Println("数据库连接失败")
		return
	}

	//sql 预处理

	stmt,err1:=db.Prepare("insert into file_info (id,file_name,file_code,file_size) values (?,?,?,?)")

	stmtQuery,queryErr:=db.Prepare("select * from file_info")

	// 查询
	rows,errQu:=db.Query()
	for rows.Next() {
		var id int
		var name string
		rows.Scan()
	}






	if queryErr != nil{
		fmt.Println("预处理error1")
		return
	}

	if err1 != nil{
		fmt.Println("预处理error1")
		return
	}


	defer func() {
		if stmt != nil{
			stmt.Close()
		}
	}()


	//不定量参数 参数和占位符对应
	result,err2:=stmt.Exec("4","sss","1213","2222")

	if err2 != nil{
		fmt.Println("预处理error2:"+err2.Error())
		return
	}


	//3。获取结果
	count,err3:=result.RowsAffected()

	if err3 != nil{
		fmt.Println("获取结果error")
		return
	}

	if count > 0 {
		fmt.Println("success")
	} else {
		fmt.Println("failure")
	}






}
