package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/api"
)

//用户主界面
//显示登录，留言等信息

func home(db *sql.DB, router *gin.Engine) *gin.Engine {
	router = api.Home(db, router)
	return router
}
