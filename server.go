package main

import "github.com/go-martini/martini"

func main() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Use(martini.Static("Flat-UI")) // 也会服务静态文件于"assets"的文件夹
	m.Run()
}
