package main

import (
	"fmt"
	"message-board/public_func"
	"message-board/service"
)

func main() {
	db, err := public_func.Getdb()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	router := service.APIs(db)
	router.Run("0.0.0.0:5000")
}
