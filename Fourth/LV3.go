package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	num, i      int
	account     string
	password    string
	words       = "傻逼"
	name        = "D:/GOPROJECTS/Homework/Fourth/users.data.txt"
	information map[string]*User
	//生成公钥和私钥
	privateKey, err = rsa.GenerateKey(rand.Reader, 1024)
	publicKey       = privateKey.PublicKey
)

func main() {
	Open()
	menu()
}

type User struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func menu() {
	for i = 1; ; {
		fmt.Println("登录注册Terminal\n1.注册\n2.修改密码\n3.退出")
		fmt.Scan(&num)
		switch num {
		case 1:
			register()
			i++ //利用i的值判断是否重复注册
		case 2:
			recode()
		default:
			return
		}
	}
}

func Open() {
	//程序启动时，从本地文件users.data读入用户名和密码并保存到map中
	file, err := os.OpenFile(name, os.O_APPEND, 0777)
	defer file.Close()
	file.Seek(0, 0)
	Error(err)

	if i > 1 {
		buf := make([]byte, 1024)
		for {
			length, _ := file.Read(buf)
			if length == 0 {
				break
			}
			content := make(map[string]interface{})
			err = json.Unmarshal(buf, &content) //将content由map转为[]byte
			Error(err)
		}
		decrypt()
	}
}
func register() {
	if i == 1 {
		fmt.Println("请输入注册账号:")
		_, err := fmt.Scan(&account)
		Error(err)
		index := strings.Index(account, words)
		//输入的账号中不能有敏感词汇
		//若符合要求，则添加进map中
		if index != -1 {
			fmt.Println("你的账号不符合要求")
			return
		}
		code()
		fmt.Println("注册成功")
	} else {
		fmt.Println("不能重复注册！")
	}
}
func code() {
	information = make(map[string]*User)
	fmt.Println("请输入注册密码：（长度大于6位）")
	fmt.Println("注意密码禁止使用+，空格，/，？，%，#，=")
	_, err := fmt.Scan(&password)
	Error(err)
	//输入的密码中限制长度且不能与账号一样
	if len(password) <= 6 || password == account {
		fmt.Println("密码格式错误")
		return
	}
	encrypt()
	//将用户名账号和密码保存的格式为json格式的文本
	save := User{Account: account, Password: password}
	information[save.Account] = &save
	information[save.Password] = &save
	data, _ := json.Marshal(information)
	err = ioutil.WriteFile("users.data", data, 0644)
	Error(err)

}
func recode() {
	information = make(map[string]*User)
	fmt.Println("请输入你修改后的新密码")
	fmt.Println("注意密码禁止使用+，空格，/，？，%，#，=")
	_, err := fmt.Scan(&password)
	if len(password) <= 6 || password == account {
		fmt.Println("密码格式错误")
		return
	}
	encrypt()
	//覆盖原来文件内容的方式写入
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	Error(err)
	n, _ := f.Seek(0, 0)
	_, err = f.WriteAt([]byte(password), n)
	reSave := User{Password: password}
	information[reSave.Password] = &reSave
	data, _ := json.Marshal(information)
	err = ioutil.WriteFile("users.data", data, 0644)
	Error(err)
	fmt.Println("修改成功")
	defer f.Close()
}
func Error(err error) {
	if err != nil {
		panic(err)
	}
}
func encrypt() {
	//加密
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte(password),
		nil)
	Error(err)
	fmt.Println("加密字节: ", encryptedBytes)
}

func decrypt() {
	//解密
	var encryptedBytes []byte
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	Error(err)
	fmt.Println("解密后的内容: ", string(decryptedBytes))
}
