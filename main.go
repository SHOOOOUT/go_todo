package main

import (
	"github.com/shuto/go_todo/db"
	"github.com/shuto/go_todo/route"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	//cssなどの静的ファイルのディレクトリを指定
	router.Static("/assets", "./assets/css")

	//dbInit()
	db.DbInit()

	//Index
	router.GET("/", route.Index)
	//Create
	router.POST("/new", route.Create)
	//Detail
	router.GET("/detail/:id", route.Detail)
	//Update
	router.POST("/update/:id", route.Update)
	//delete確認
	router.GET("/delete_check/:id", route.DelConf)
	//Delete
	router.POST("/delete/:id", route.Delete)

	router.Run()

}
