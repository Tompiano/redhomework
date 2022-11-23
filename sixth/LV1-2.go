package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	id       int
	username string
	phone    string
	password string
	question string
	answer   string
}

var (
	user User
	db   *sql.DB
)

func main() {
	test()
	r := gin.Default()
	r.Use(HandlerFunc()) //设置全局中间件
	//设置cookie保持用户状态
	r.GET("/cookie", SetCookie)
	//注册
	r.POST("/register", register)
	//登录
	r.GET("/login", login)
	//忘记密码
	r.GET("/forget", Forget)

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
func HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("中间件")
	}
}
func Error(err error) {
	if err != nil {
		log.Printf("insert data error : %v", err)
		return
	}
}
func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

	}
	fmt.Println(cookie)
}
func login(c *gin.Context) {
	username := c.Query("username") //获取用户名
	password := c.Query("password") //获取用户密码
	stmt, err := db.Prepare("select * from information where username=?")
	Error(err)
	row, err := stmt.Query(username)
	Error(err)
	defer row.Close() //延迟关闭
	for row.Next() {
		err = row.Scan(&user.username, &user.password, &user.phone)
		Error(err)
	}
	if username != user.username {
		c.JSON(200, "该用户不存在")
	} else {
		if password != user.password {
			c.JSON(200, "用户名与密码不匹配")
		} else {
			c.JSON(200, "登录成功")
		}
	}
}
func register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	phone := c.Query("phone")
	//获取用户设置的保密问题以及答案
	question := c.Query("question")
	answer := c.Query("answer")
	if question == "" || answer == "" {
		c.JSON(200, "问题和答案都必须要填")
	}
	//插入和用户信息有关的数据进数据库
	result, err := db.Exec("insert into information(username,phone,password,question,answer) value (?,?,?,?,?)",
		username, phone, password, question, answer)
	Error(err)
	// 返回新插入数据的id
	result.LastInsertId()
	// 返回影响的行数
	result.RowsAffected()

	c.JSON(200, "注册成功")
}
func Forget(c *gin.Context) {
	name := c.Query("username")
	question := c.Query("question")
	answer := c.Query("answer")        //获取用户对保密问题的答案
	newPassword := c.Query("password") //获取修改后的密码
	stmt, err := db.Prepare("select * from information where username=?")
	Error(err)
	row, err := stmt.Query(name)
	Error(err)
	defer row.Close() //延迟关闭
	for row.Next() {  //从数据库中获取用户的相关信息
		err = row.Scan(&user.username, &user.password, &user.phone)
		Error(err)
	}
	if name != user.username {
		c.JSON(200, "该用户不存在")
	} else if question == user.question && answer == user.answer {
		c.JSON(200, "认证成功")
		result, err := db.Exec("update information set password=?", newPassword)
		Error(err)
		result.LastInsertId()
		result.RowsAffected()
	}
	c.JSON(200, "成功找回密码")
}
