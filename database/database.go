package database

import (
	"database/sql"
	"taskManagerUserService/model"
	"fmt"
)

func CreateUser(db *sql.DB, user *model.User) error {
	fmt.Println(user.UserName)
	fmt.Println(user.EmailId)
	fmt.Println(user.Password)
	_, queryErr := db.Query(`INSERT INTO task_manager_user ("userName","emailId","password") VALUES($1,$2,$3)`, user.UserName, user.EmailId, user.Password);
	if (queryErr != nil) {
		fmt.Println(queryErr)
		return queryErr
	}
	return nil
}
