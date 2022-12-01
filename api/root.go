package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

// 主页的api，get
func root(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cuname, err1 := c.Cookie("username")
		//只有当cookie username和password同时存在时
		resp, err := public_func.Cookie_check(db, cuname, err1)
		//当cookie有误时，返回json
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": resp})
		} else {
			c.JSON(200, map[string]interface{}{"ok": true, "data": "你好，用户" + cuname})
		}

	}
}
