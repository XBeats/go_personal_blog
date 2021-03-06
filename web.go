package main

import (
	"go_personal_blog/controller"
	"go_personal_blog/controller/common"
	"go_personal_blog/core"
)

func main() {
	core.Router("/", &controller.MainController{})
	core.Router("/m", &controller.MarkDown{})
	core.Router("/markdown", &controller.MarkDown{})
	core.Router("/welcome", &controller.Welcome{})
	core.Router("/json1", &common.Common{})
	core.Router("/t", &controller.MarkTest{})
	json := &controller.JsonController{}
	core.Router("/json", json)
	core.Router("/j", json)

	core.SetStaticPath("/css", "static/css")
	core.SetStaticPath("/font", "static/font")
	core.SetStaticPath("/icon", "static/icon")
	core.SetStaticPath("/js", "static/js")
	core.SetStaticPath("/images", "static/img")
	core.Run()
}
