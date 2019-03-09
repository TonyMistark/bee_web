package main

import (
    "github.com/astaxie/beego"
    "time"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    time.Sleep(time.Second * 1)
    this.Ctx.WriteString("hello world")
}

func main() {
    beego.Router("/hello/world/$", &MainController{})
    beego.Run()
}
