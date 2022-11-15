package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	userregister UserRegister
)

func main() {
	route := gin.Default()
	route.Use(HandlerFunc())
	user := route.Group("/user")
	{
		user.GET("/cookie", SetCookie) //设置cookie
		user.POST("/register", Reg)    //注册
		user.GET("/login", Login)      //登录
	}
	route.Run()
}
func Reg(c *gin.Context) {
	//绑定用户的姓名及密码
	err := c.ShouldBind(&userregister)
	if err != nil {
		fmt.Println("注册失败")
		log.Fatal(err.Error())
		return
	}
	//用map保存用户信息
	data["Username"] = userregister.Username
	data["Password"] = userregister.Password
	//如果绑定有值的话就返回
	c.Writer.Write([]byte("注册成功"))
}
func Login(c *gin.Context) {
	//判断用户是否存在
	_, Judge := data["Username"]
	if !Judge {
		c.Writer.Write([]byte("你还未注册"))
	} else {
		c.Writer.Write([]byte("登录成功"))
	}
	return
}
func SetCookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

	}
	fmt.Println(cookie)
}
func HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("一个中间件")
	}
}
