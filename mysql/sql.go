package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.2.md

func main() {
	db, err := sql.Open("mysql", "username:passwd@tcp(hostname:port)/dbname?charset=utf8")
	checkErr(err)
    
    //for connect pool
    db.SetMaxOpenConns(2000)
    db.SetMaxIdleConns(1000)
    db.Ping()
   
    /*
	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	
    //删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

    */
   
	//查询数据
	rows, err := db.Query("SELECT id,name,user,mem FROM tablename where name='name'")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var user string
		var mem string
		err = rows.Scan(&id, &name, &user, &mem)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(user)
		fmt.Println(mem)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

