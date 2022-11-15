package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//设置Cookie,保持用户状态
	router.GET("/cookie", Set)
	//有cookie的时候和无cookie的时候返回不同的结果
	router.GET("/login", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			fmt.Println(cookie)
		} else {
			fmt.Println("您还未登录！")
		}
	})

	router.Run()
}
func Set(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)

	} else {
		fmt.Println("hhh")
	}
	fmt.Println(cookie)
}
