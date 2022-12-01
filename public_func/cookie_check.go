package public_func

import "database/sql"

func Cookie_check(db *sql.DB, username string, err1 error) (map[string]interface{}, error) {
	//第一种情况，没有cookie时
	if err1 != nil {
		//如果cookie为空，返回错误信息
		return map[string]interface{}{"ok": false, "data": "no cookie"}, err1
	} else { //第二种情况：提供了cookie
		datalist, err := SelectFromUserName(db, username)
		//当查不到此用户，说明cookie有误
		if err != nil {
			return map[string]interface{}{"ok": false, "data": "no such username"}, err
		}
		//提取列表的第一条，这就是查询的信息
		data := datalist[0]
		//校验成功
		if data.Name == username {
			return map[string]interface{}{"ok": true, "data": "check_success"}, nil
		}
	} //默认返回
	return nil, err1
}
