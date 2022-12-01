package public_func

import (
	"database/sql"
	"log"
)

func Insert_data(db *sql.DB, username string, password string, question string, question_content string) bool {
	_, err := db.Exec("insert into user (username ,password ,question, question_content) value (?,?,?,?)", username, password, question, question_content)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
