package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

// APIs api的入口处

func Home(db *sql.DB, router *gin.Engine) *gin.Engine {
	router.GET("/", root(db), func(c *gin.Context) {
	})
	router.POST("/register", regis(db), func(c *gin.Context) {
	})
	router.GET("/login", login(db), func(c *gin.Context) {
	})
	router.GET("/forget", forget(db), func(c *gin.Context) {
	})
	//相同地址，获取用户密保问题，提供给前端调用，data中保存密保问题
	router.GET("/forcheck", forCheck(db), func(c *gin.Context) {
	})

	return router
}

func Message_board(db *sql.DB, router *gin.Engine) *gin.Engine {

	router.POST("/leave_message", leave_message(db), func(c *gin.Context) {
	})
	router.GET("/get_table", get_table(db), func(c *gin.Context) {
	})
	router.PUT("/del_message", del_message(db), func(c *gin.Context) {
	})
	router.PUT("/alter_message", alter_comment(db), func(c *gin.Context) {
	})
	return router
}
