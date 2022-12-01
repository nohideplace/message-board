package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func APIs(db *sql.DB) *gin.Engine {

	router := gin.Default()
	router = home(db, router)
	router = message_board(db, router)
	//用户留言，传入父id，用户名，内容
	return router
}
