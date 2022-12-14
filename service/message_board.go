package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"message-board/api"
)

//用户留言界面
//将api接口封装后放在这里

//登录后使用功能：

/*
一.留言功能
1.查看留言
2.发表留言
3.修改自己的留言
二.评论功能
1.查看（自己的）评论
2.发表评论
3.删除评论

三.实现方式，新建一张评论表，在用户留言时加一条记录，标识它的评论者，以及它的父评论
列出评论列表：先列出它自己的评论，然后去数据库中查找父评论为它的id的评论，就是它的子评论，然后这些评论进行递归的相同操作，直到没有父评论为它
*/
func message_board(db *sql.DB, router *gin.Engine) *gin.Engine {
	router = api.Message_board(db, router)
	return router
}
