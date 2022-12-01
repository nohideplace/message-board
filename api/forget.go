package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

func forget(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.DefaultQuery("username", "null")
		if username == "null" {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "lack of content"})
		}
		datalist, err := public_func.SelectFromUserName(db, username)
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "get data wrong"})
		} else { //存在则返回用户密保
			data := datalist[0]
			c.JSON(200, map[string]interface{}{"ok": true, "data": data.Question})
		}
	}
}
