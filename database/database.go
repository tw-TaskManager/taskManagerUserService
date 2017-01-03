package database

import (
	"database/sql"
	"taskManagerUserService/model"
	"fmt"
)

func CreateUser(db *sql.DB, user *model.User) error {
	_, queryErr := db.Query(`INSERT INTO task_manager_user ("userName","emailId","password") VALUES($1,$2,$3)`, user.UserName, user.EmailId, user.Password);
	if (queryErr != nil) {
		fmt.Println(queryErr)
		return queryErr
	}
	return nil
}

func Login(db *sql.DB, user *model.User) (model.User, error) {
	row, queryErr := db.Query(`SELECT "userName","password","userId" from task_manager_user where "emailId"=$1`, user.EmailId);
	if (queryErr != nil) {
		fmt.Println(queryErr)
		return model.User{}, queryErr
	}
	userInfo := make([]model.User, 0, 0)
	for row.Next() {
		var userName string
		var password string
		var userId int32
		row.Scan(&userName, &password,&userId)
		user := model.User{}
		user.UserName = userName
		user.Password = password
		user.UserId = userId
		userInfo = append(userInfo, user)
	}
	if (len(userInfo) == 0) {
		return model.User{}, nil
	}
	return userInfo[0], nil

}
