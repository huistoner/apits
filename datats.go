package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type countInfo struct {
	ip          string
	mac         string
	channelName string
	starTime    string
	endTime     string
	enter       int32
	leave       int32 //xml中使用exit表示离开人数，这里改用leave表示。
}

func main() {
	var uername = "sto"
	var passwd = "@2021MP3abc"
	var zhuji = "cd-cdb-99uyhax8.sql.tencentcdb.com"
	var port = "63118"
	var data = "counting"
	var dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", uername, passwd, zhuji, port, data)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("链接数据库失败，", err)
	} else {
		stmt, _ := db.Prepare("select top1 ip, mac,channelName,starTime,endTime,enter,`leave` from people")
		rows, _ := stmt.Query()
		defer rows.Close()
	}
}