package main

import (
	"fmt"
	"message-board/api"
	"message-board/public_func"
)

func main() {
	db, err := public_func.Getdb()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	router := api.APIs(db)
	router.Run("0.0.0.0:5000")
}
