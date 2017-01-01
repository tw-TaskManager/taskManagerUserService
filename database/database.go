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

func Login(db *sql.DB, user *model.User) (string, error) {
	row, queryErr := db.Query(`SELECT "emailId" from task_manager_user where "userName"=$1 and "password"=$2`, user.UserName, user.Password);
	if (queryErr != nil) {
		fmt.Println(queryErr)
		return "", queryErr
	}
	userEmail := make([]string, 0, 0)
	for row.Next() {
		var emailId string
		row.Scan(&emailId)
		userEmail = append(userEmail, emailId)
	}
	if(len(userEmail)==0) {
		return "",nil
	}
	return userEmail[0], nil

}
