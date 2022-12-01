package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

func cookie_check(db *sql.DB, username string, err1 error) (map[string]interface{}, error) {
	//第一种情况，没有cookie时
	if err1 != nil {
		//如果cookie为空，返回错误信息
		return map[string]interface{}{"ok": false, "data": "no cookie"}, err1
	} else { //第二种情况：提供了cookie
		datalist, err := public_func.SelectFromUserName(db, username)
		//当查不到此用户，说明cookie有误
		if err != nil {
			return map[string]interface{}{"ok": false, "data": "no such username"}, err
		}
		//提取列表的第一条，这就是查询的信息
		data := datalist[0]
		//校验成功
		if data.Name == username {
			return map[string]interface{}{"ok": true, "data": "check_success"}, nil
		}
	} //默认返回
	return nil, err1
}

// APIs api的入口处
func APIs(db *sql.DB) *gin.Engine {

	router := gin.Default()
	router.GET("/", root(db), func(c *gin.Context) {
	})
	router.POST("/register", regis(db), func(c *gin.Context) {
	})
	router.GET("/login", login(db), func(c *gin.Context) {
	})
	router.POST("/forget", forget(db), func(c *gin.Context) {
	})
	//相同地址，获取用户密保问题，提供给前端调用，data中保存密保问题
	router.GET("/forget", forCheck(db), func(c *gin.Context) {
	})
	//用户留言，传入父id，用户名，内容
	router.POST("/leave_message", leave_message(db), func(c *gin.Context) {
	})
	router.GET("/get_table", get_table(db), func(c *gin.Context) {
	})
	router.POST("/del_message", del_message(db), func(c *gin.Context) {
	})
	router.POST("/alter_message", alter_comment(db), func(c *gin.Context) {
	})
	return router
}
