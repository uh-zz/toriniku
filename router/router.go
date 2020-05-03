package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/me/toriniku/controllers"
	// "github.com/me/toriniku/db"
)

func Router(dbConn *gorm.DB) {
	nikuHandler := controllers.NikuHandler{
		Db: dbConn,
	}
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/top", nikuHandler.GetAll) // 一覧画面
	// r.GET("/json", nikuHandler.Getjson) // json画面
	// r.POST("/todo", nikuHandler.CreateTask)            // 新規作成
	// r.GET("/todo/:id", nikuHandler.EditTask)           // 編集画面
	// r.POST("/todo/edit/:id", nikuHandler.UpdateTask)   // 更新
	// r.POST("/todo/delete/:id", nikuHandler.DeleteTask) // 削除
	r.Run(":8000")
}
