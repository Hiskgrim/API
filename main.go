package main

import (
	_ "cdve_api_crud/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:1234@127.0.0.1:5434/udistrital?sslmode=disable&search_path=argo,financiera,agora,core,cdve")
}

func main() {
	orm.Debug = true
	if beego.BConfig.RunMode == "dev" {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	AllowHeaders: []string{"Origin", "x-requested-with",
	"content-type",
	"accept",
	"origin",
	"authorization",
	"x-csrftoken"},
	ExposeHeaders: []string{"Content-Length"},
	AllowCredentials: true,
	}))
	beego.Run()
}
