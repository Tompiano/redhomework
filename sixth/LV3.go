package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	userid   int
	username string
	phone    string
	password string
	msg      string
}

var (
	u  User
	db *sql.DB
)

/*
我想的是除了创建的装用户信息的表之外，再创建一张表用于装用户的留言
根据不同的username来识别不同的用户，不同的id来识别留言的先后，相当于盖的楼层编号（我觉得。。。）
但是吧，我觉得每条留言本身就对应了不同的id。。。
所以查看留言表的时候我就只设了id这一个查询条件。。。
另外对于实现回复，我觉得在他说不要求嵌套的条件下要实现的话，就直接在下面接着写留言应该就算是回复吧。。
（理解能力不太行，抱紧狗头）
*/
func main() {
	test()
	r := gin.Default()
	r.Use(HandlerFunc())        //设置全局中间件
	r.GET("/cookie", SetCookie) //设置cookie保持用户状态
	r.GET("/login", login)      //登录
	r.GET("/look", look)        //查看留言
	r.POST("write", write)      //写留言
	r.Run()
}
func test() {
	var dns = "root:123456@tcp(127.0.0.1:3306)/user" //DSN（数据源名称）
	var err error
	db, err = sql.Open("mysql", dns) //连接数据库
	err = db.Ping()                  //检查数据库是否可用且可访问
	if err != nil {
		fmt.Println("数据库连接失败")
		log.Println(err)
	}
}
func login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	stmt, err := db.Prepare("select * from information where username=?")
	Error(err)
	row, err := stmt.Query(username)
	Error(err)

	defer row.Close()
	for row.Next() {
		err = row.Scan(&u.username, &u.userid, &u.password)
		Error(err)
	}
	if username != u.username {
		c.JSON(200, "该用户不存在")
	} else if password != u.password {
		c.JSON(200, "你的密码不正确")
	} else {
		c.JSON(200, "登录成功")
	}

}
func HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("中间件")
	}
}
func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("cookie")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie("cookie", "test", 3600, "/", "localhost", false, true)
	}
	fmt.Println(cookie)
}
func write(c *gin.Context) {
	//获取用户写的留言内容
	message := c.PostForm("message")
	//对留言内容格式进行判断，符合格式的才能发表
	if message == "" {
		c.JSON(200, "留言不能为空")
	} else {
		msg := c.PostForm("msg")
		username := c.PostForm("username") //获取用户自己的或者是用户想要留言的用户的名字
		//通过用户名在对应的表中插入留言信息
		var result, err = db.Exec("insert into message(msg) values (?) where username=?", msg, username)
		Error(err)
		c.JSON(200, "成功留下留言")
		// 返回新插入数据的id
		_, err = result.LastInsertId()
		Error(err)
		// 返回影响的行数
		_, err = result.RowsAffected()
		Error(err)
	}
}
func look(c *gin.Context) {
	id := c.PostForm("userid")
	//查找id对应的留言内容
	stmt, err := db.Prepare("select * from message where id=?")
	Error(err)
	row, err := stmt.Query(id)
	Error(err)
	defer row.Close()
	for row.Next() {
		err = row.Scan(&u.userid, &u.username, &u.msg)
		c.JSON(200, "6")
		Error(err)
	}
	c.JSON(200, "查看成功")
}
func Error(err error) {
	if err != nil {
		log.Println("错误：", err)
		return
	}
}
