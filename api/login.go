package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

func check_user_pass(db *sql.DB, username string, password string) bool {
	datalist, err := public_func.SelectFromUserName(db, username)
	if err != nil {
		return false
	}
	if datalist == nil {
		return false
	}
	data := datalist[0]
	if password == data.Password {
		return true
	}
	return false
}

func login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.DefaultQuery("username", "null")
		password := c.DefaultQuery("password", "null")
		if username == "null" || password == "null" {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "no username or password"})
		}
		ok := check_user_pass(db, username, password)
		if ok {
			c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
			c.JSON(200, map[string]interface{}{"ok": true, "data": "用户登录成功，接下来返回主页"})
		} else {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "登录失败，用户不存在或密码错误"})
		}
	}
}
