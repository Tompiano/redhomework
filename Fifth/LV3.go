package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

type Struct struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	ID       string `json:"id"`
}

var (
	str  Struct
	data = make(map[string]string)
	path = "D:/GOPROJECTS/text.txt"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(HandlerFunc())
	//用追加模式打开文件并通过文件读取用户信息
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	Error(err)
	_, err = file.Seek(0, 0)
	Error(err)
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		m, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if m == 0 {
			break
		}
	}
	//用户注册
	r.POST("/bind", func(c *gin.Context) {
		err = c.ShouldBind(&str)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		//将用户信息保存在map中
		data["Name"] = "name"
		data["Password"] = "password"
		data["ID"] = "id"
		//返回
		c.Writer.Write([]byte("绑定用户信息"))
	})
	//用户注册后的信息追加到文件中
	write := bufio.NewWriter(file) //使用带缓存的 *Writer将用户的信息写入文件
	for i := 0; i < len(data); i++ {
		write.WriteString(data["Name"])
		write.WriteString(data["Password"])
		write.WriteString(data["ID"])
	}
	write.Flush()
	r.Run()
}
func HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("一个中间件")
	}
}
func Error(err error) {
	if err != nil {
		panic(err)
		return
	}
}
