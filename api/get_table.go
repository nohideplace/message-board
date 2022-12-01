package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

// 编写一个遍历评论区的函数，返回结构体切片的切片，保存了评论区的所有内容
type MessageForm struct {
	Id       int    `form:"id"`
	ParentId int    `form:"parentid"`
	Name     string `form:"name"`
	Message  string `form:"message"`
	IsDelete int    `form:"isdelete"`
}

func SelectTableFromParentId(db *sql.DB, id int) ([]MessageForm, error) {
	var data MessageForm
	var list []MessageForm
	rows, err := db.Query("select * from message where parentid=?", id)
	if err != nil {
		return []MessageForm{}, err
	}
	defer rows.Close()
	for rows.Next() {
		// row.scan 必须按照先后顺序 &获取数据
		//var temp interface{}
		//由于表的结构是固定的，这里可以写死格式
		//先后顺序由表的结构决定
		err := rows.Scan(&data.Id, &data.Name, &data.ParentId, &data.Message, &data.IsDelete)
		if err != nil {
			log.Println(err)
			return []MessageForm{}, err
		}
		list = append(list, data)
	}
	return list, nil
}
func SelectTableFromId(db *sql.DB, id int) ([]MessageForm, error) {
	var data MessageForm
	var list []MessageForm
	rows, err := db.Query("select * from message where id=?", id)
	if err != nil {
		return []MessageForm{}, err
	}
	defer rows.Close()
	for rows.Next() {
		// row.scan 必须按照先后顺序 &获取数据
		//var temp interface{}
		//由于表的结构是固定的，这里可以写死格式
		err := rows.Scan(&data.Id, &data.Name, &data.ParentId, &data.Message, &data.IsDelete)
		if err != nil {
			log.Println(err)
			return []MessageForm{}, err
		}
		list = append(list, data)
	}
	return list, nil
}

func get_from_parentid(db *sql.DB, id int, dataform []MessageForm) []MessageForm {
	//查到第一层的数据，继续向下查找
	//先获取到所有父id为该id的节点列表
	//这些都是同级
	datalist, err := SelectTableFromParentId(db, id)
	//查找报错
	if err != nil {
		return nil
	}

	//上面只能说明查找没有报错
	//如果数组长度为0，说明不存在了，结束递归，将前面积累的信息返回
	if len(datalist) == 0 {
		return dataform
	}
	for _, i := range datalist {
		//遍历当前层的所有元素的id，向下查找元素
		//先将这一层的所有数据append进入切片
		for _, j := range datalist {
			dataform = append(dataform, j)
		}

		return get_from_parentid(db, i.Id, dataform)
	}

	return nil
}

// 获取所有root节点，返回一个root节点的id列表，用于后面遍历
func get_all_root(db *sql.DB) (resp []int, err error) {
	datalist, err := SelectTableFromParentId(db, 114514)
	if err != nil {
		return nil, err
	}
	for _, i := range datalist {
		resp = append(resp, i.Id)
	}
	return resp, nil
}

//func test(db *sql.DB) string {
//	var table [][]MessageForm
//	rootlist, err := get_all_root(db)
//	if err != nil {
//		return "获取根列表失败"
//	}
//
//}

func get_table(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var table [][]MessageForm
		rootlist, err := get_all_root(db)
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "get table fail"})
		}
		for _, i := range rootlist {
			var datalist []MessageForm
			self, err := SelectTableFromId(db, i)
			if err != nil {
				c.JSON(200, map[string]interface{}{"ok": false, "data": "get table fail"})
			}
			//首先在每次都先把它本身加入数组
			//id搜索是唯一的，直接取0
			datalist = append(datalist, self[0])
			datalist = get_from_parentid(db, i, datalist)
			if datalist == nil {
				c.JSON(200, map[string]interface{}{"ok": false, "data": "datalist nil!"})
			}
			table = append(table, datalist)
		}
		//data, err := json.MarshalIndent(table, "", "  ")
		//fmt.Println(string(data))
		if err != nil {
			c.JSON(200, map[string]interface{}{"ok": false, "data": "get table fail"})

		} else {
			c.JSON(200, map[string]interface{}{"ok": true, "data": table})
		}

	}
}
