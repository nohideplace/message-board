package public_func

import (
	"database/sql"
	"log"
)

type user struct {
	Id              int
	Name            string
	Password        string
	Question        string
	QuestionContent string
}

func SelectFromUserName(db *sql.DB, username string) ([]user, error) {
	var data user
	var list []user
	rows, err := db.Query("select * from user where username=?", username)
	if err != nil {
		return []user{}, err
	}
	defer rows.Close()
	for rows.Next() {
		// row.scan 必须按照先后顺序 &获取数据
		//var temp interface{}
		//由于表的结构是固定的，这里可以写死格式
		err := rows.Scan(&data.Id, &data.Name, &data.Password, &data.Question, &data.QuestionContent)
		if err != nil {
			log.Println(err)
			return []user{}, err
		}
		list = append(list, data)
	}
	return list, nil
}
