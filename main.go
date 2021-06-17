package main

import (
	_ "golang-project-prototype/config"
	_ "golang-project-prototype/library/util/rdc"
	_ "golang-project-prototype/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
