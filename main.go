package main

import (
	_ "demo/routers"

	// 这是添加的注释
	"github.com/astaxie/beego"
)

func main() {

	beego.Run()
}
