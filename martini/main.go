package main

import (
	"github.com/go-martini/martini"
	"fmt"
)

func main() {
	fmt.Println(fmt.Sprintf(`(?P<%s>[^/#?]+)`, "hello world"))
	m := martini.Classic()//创建一个典型的martini实例
	m.Get("/", func() string {     //接收对\的GET方法请求，第二个参数是对一请求的处理方法
		return "Hello world!"
	})
	m.Run()//运行服务器

}