package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

//删除留言，注意不是将它的数据给删除了，而是给它打一个标记，让其内容不被显示（该留言已被删除），但是它的所有子评论还是要显示

// 使用时鉴权，当用户为它的拥有者时，可以执行某条信息的删除操作
// 删除时，将该条信息标记为已删除，在这个构造返回数据
func update_isdelete(db *sql.DB, message_id int) bool {
	//update <表名> set <字段> =<值> [, <字段> =<值>...] [where 表达式]
	_, err := db.Exec("update message set isdelete=1 where id=?", message_id)
	if err != nil {
		return false
	}
	return true
}

func del_message(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cuname, err1 := c.Cookie("username")
		//只有当cookie username和password同时存在时
		resp, err := cookie_check(db, cuname, err1)
		//当cookie有误时，返回json
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": resp})
		}
		var message MessageForm
		if err3 := c.ShouldBind(&message); err3 != nil {
			// 处理错误请求
			c.JSON(200, map[string]interface{}{"ok": false, "data": "bind data wrong"})
		}
		ok := update_isdelete(db, message.Id)
		if ok {
			c.JSON(200, map[string]interface{}{"ok": true, "data": "delete successfully"})
		} else {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "sql wrong"})
		}

	}

}
