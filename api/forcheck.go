package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

func forCheck(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.DefaultQuery("username", "null")
		questionContent := c.DefaultQuery("question_content", "null")
		if username == "null" || questionContent == "null" {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "lack of content"})
		}
		datalist, err := public_func.SelectFromUserName(db, username)
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "get data wrong"})
		}
		data := datalist[0]
		if data.QuestionContent == questionContent {
			c.JSON(200, map[string]interface{}{"ok": true, "data": "你的密码是" + data.Password + "\n下次别再忘了哦"})
		} else {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "密保问题错误啦"})
		}

	}
}
