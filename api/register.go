package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

func register(db *sql.DB, username string, password string, question string, question_content string) (bool, string) {
	//当不存在时，创建用户
	ok := public_func.Insert_data(db, username, password, question, question_content)
	if !ok {
		return ok, "插入失败，可能已经存在该用户"
	}
	return true, "插入成功"
}
func regis(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.DefaultQuery("username", "null")
		password := c.DefaultQuery("password", "null")
		question := c.DefaultQuery("question", "null")
		question_content := c.DefaultQuery("question_content", "null")
		//当用户名和密码有任何一部分没有被提供时，返回信息
		if username == "null" || password == "null" || question == "null" || question_content == "null" {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "lack of content"})
		}
		//当全部内容都非空时，注册用户
		ok, _ := register(db, username, password, question, question_content)
		if ok {
			//设置cookie在根
			c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
			c.JSON(200, map[string]interface{}{"ok": true, "data": "用户创建成功，接下来返回主页"})
		} else {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "用户创建失败，也许用户已经存在了罢"})
		}
	}
}
