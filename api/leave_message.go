package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

/*
一.业务流程：
1.所有用户可以进入评论区，查看所有用户的发言
2.用户可以向评论区发言
(1)发言时附带评论id，则将该评论的id作为自己的父id
(2)发言时不带评论id，传入114514，作为一条新的根评论，父评论id设置为114514
3.遍历评论区，
(1)首先找到所有父id为114514的评论，从它们开始往下递归，每个这样的评论都存在一个切片内
(2)递归在数据表中查找所有父id为它的评论，每找到一个就往切片中append一个元素，最后遍历完毕后，将这个切片append到(1)中的切片内
(3)最后返回一个装满所有评论的结构体切片的切片，还有执行成功或失败的返回
*/

//type MessageForm struct {
//	Id       int    `form:"id"`
//	ParentId int    `form:"parentid"`
//	Name     string `form:"name"`
//	Message  string `form:"message"`
//	IsDelete int    `form:"isdelete"`
//}

//数据表的构造
//CREATE TABLE `message` (
//`id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
//`username` varCHAR(16) NOT NULL,
//`parentid` int NOT NULL,
//`message`  varchar(1000) NOT NULL default "no message",
//`isdelete` int default 0
//);

// 留言，插入留言的id，留言用户，留言内容，父评论id
func insert_message(db *sql.DB, username string, content string, parentid int) bool {
	_, err := db.Exec("insert into message (username, message, parentid) value (?,?,?)", username, content, parentid)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func leave_message(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		cuname, err1 := c.Cookie("username")
		//只有当cookie username和password同时存在时
		resp, err := cookie_check(db, cuname, err1)
		//当cookie有误时，返回json
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": resp})
		}
		var messageform MessageForm
		if err3 := c.ShouldBind(&messageform); err3 != nil {
			// 处理错误请求
			c.JSON(200, map[string]interface{}{"ok": false, "data": "bind data wrong"})
		}
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "no such user"})
		} else {
			ok := insert_message(db, messageform.Name, messageform.Message, messageform.ParentId)
			if ok {
				c.JSON(200, map[string]interface{}{"ok": true, "data": "插入评论成功"})
			} else {
				c.JSON(200, map[string]interface{}{"ok": false, "data": "插入评论失败"})
			}
		}
		//当cookie无误时，

	}
}
