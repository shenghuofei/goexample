package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

//https://beego.me/docs/mvc/model/orm.md

type User struct {
    Id   int
    Username string `orm:"size(100)"`
    Email string `orm:"size(100)"`
    Password string `orm:"size(100)"`
}

func init() {
    // 设置默认数据库
    orm.RegisterDataBase("default", "mysql", "username:passwd@/dbname?charset=utf8", 30)
    
    orm.SetMaxIdleConns("default", 30)
    orm.SetMaxOpenConns("default", 30)
    orm.Debug = true

    // 注册定义的 model
    orm.RegisterModel(new(User))
	//RegisterModel 也可以同时注册多个 model
	//orm.RegisterModel(new(User), new(Profile), new(Post))

    // 创建 table
    orm.RunSyncdb("default", false, true)
}


func main(){
    var o orm.Ormer
	var rs orm.RawSeter
	o = orm.NewOrm()
   
    name := "test"
    user := User{Username:name}
    
    // 插入表one
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
   
    // 更新表one
    user.Username = "test1"
    //num, err := o.Update(&user) //更新所有字段
    num,err := o.Update(&user, "Username") //更新指定字段
    fmt.Printf("NUM: %d, ERR: %v\n", num, err) 
  
    // 读取 one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Printf("ERR: %v %v\n", err,u)

    // 删除one
    num, err = o.Delete(&u)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
 
    //插入多个
    users := []User{
    {Username: "slene"},
    {Username: "astaxie"},
    {Username: "unknown"},
    }
    
    successNums, err := o.InsertMulti(100, users) //第一个参数是并发度
    fmt.Printf("snum:%d ERR:%v\n",successNums,err)
    
    var rusers []User
    qs := o.QueryTable(user)
    qs.Filter("Username", "test1").Exclude("Id", 9).All(&rusers)
	// WHERE profile.age IN (18, 20) AND NOT profile_id < 1000
    //num,err = qs.OrderBy("id", "-username").All(&rusers)
	// ORDER BY id ASC, username DESC
     num,err = qs.OrderBy("id", "-username").Filter("Username", "test1").Exclude("Id", 9).All(&rusers)
    fmt.Printf("rnum:%d ERR:%v res:%v\n",num,err,rusers)

    //原生sql使用
	rs = o.Raw("SELECT * FROM user "+
		"WHERE username=? AND id<10 "+
		"ORDER BY id DESC "+
		"LIMIT 100", "test")
	var user_list []User
	num, err = rs.QueryRows(&user_list)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
		//fmt.Println(user_list[0].Id)
		fmt.Println(user_list)
	}
}
