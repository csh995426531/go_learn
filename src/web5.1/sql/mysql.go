package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"net/http"
)

func main() {
	http.HandleFunc("/user", doUser)
	err := http.ListenAndServe(":9090", nil)
	checkErr(err)
}

func doUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	db, err := sql.Open("mysql", "root:asd123@/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname='?',created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("insert id:", id)

	//更新数据
	//stmt, err = db.Prepare("UPDATE userinfo SET username=?,departname=?,created=? WHERE uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec("cuisaihang", "形成部门", "2018-01-01", id)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println("update affect:", affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string

		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println("uid:", uid)
		fmt.Println("username:", username)
		fmt.Println("departname:", departname)
		fmt.Println("created:", created)
	}

	//删除数据
	//stmt, err = db.Prepare("DELETE FROM userinfo WHERE uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println("delete affect:", affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
