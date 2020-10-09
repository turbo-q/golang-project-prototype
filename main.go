package main

import (
	_ "recitationSquare/config"
	_ "recitationSquare/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
