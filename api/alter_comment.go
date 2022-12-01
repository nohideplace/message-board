package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/public_func"
)

// 修改一套评论的内容
func update_comment(db *sql.DB, content string, message_id int) bool {
	//update <表名> set <字段> =<值> [, <字段> =<值>...] [where 表达式]
	_, err := db.Exec("update message set message=? where id=?", content, message_id)
	if err != nil {
		return false
	}
	return true
}

// 修改评论内容，传入message内容和id
func alter_comment(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cuname, err1 := c.Cookie("username")
		//只有当cookie username和password同时存在时
		resp, err := public_func.Cookie_check(db, cuname, err1)
		//当cookie有误时，返回json
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": resp})
		}
		var message MessageForm
		if err3 := c.ShouldBind(&message); err3 != nil {
			// 处理错误请求
			c.JSON(200, map[string]interface{}{"ok": false, "data": "bind data wrong"})
		}
		ok := update_comment(db, message.Message, message.Id)
		if ok {
			c.JSON(200, map[string]interface{}{"ok": true, "data": "alter successfully"})
		} else {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "sql wrong"})
		}

	}

}
