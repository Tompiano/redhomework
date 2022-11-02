package main

import (
	"fmt"
	"io"
	"os"
)

var (
	err     error
	file    *os.File
	name    = "D:/GOPROJECTS/homework/third/plan.txt"
	str     = "I’m not afraid of difficulties and insist on learning programming"
	content = []byte(str)
	con     []byte
)

func main() {
	//创建文件
	file, err = os.Create(name)
	Error(err)

	//用追加模式打开文件
	file, err = os.OpenFile(name, os.O_APPEND, 0666)
	defer file.Close()

	//写文件
	n, err := file.Write(content)
	Error(err)
	fmt.Println(n)

	//读取文件
	file, _ = os.Open(name)
	file.Seek(0, 0)
	buf := make([]byte, 1024)
	for {
		m, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("err=", err)
			return
		}
		if m == 0 {
			break
		}
		con = append(con, content[:]...)
	}

	fmt.Println(string(con))
	file.Close()

}
func Error(err error) {
	if err != nil {
		fmt.Println("error")
		return
	}
}
