package handler

import (
	"database/sql"
	"net/http"
	_"github.com/lib/pq"
	"io/ioutil"
	"log"
	"taskManagerClient/contract"
	"github.com/golang/protobuf/proto"
	"taskManagerUserService/model"
	"taskManagerUserService/database"
	"time"
)

func CreateUserTask(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		body, err := ioutil.ReadAll(req.Body);
		if (err != nil) {
			log.Fatalln("got error while reading request")
			return;
		}

		userInfo := &contract.User{}
		err = proto.Unmarshal(body, userInfo)
		if (err != nil) {
			log.Fatalln("got error while unmarsling")
			return;
		}

		user := model.User{}
		user.UserName = *userInfo.UserName
		user.EmailId = *userInfo.EmailId
		user.Password = *userInfo.Password
		err = database.CreateUser(db, &user);
		if (err != nil) {
			res.WriteHeader(http.StatusConflict)
			return
		}

		return

	}
}

func LoginUser(db *sql.DB) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		body, err := ioutil.ReadAll(req.Body)
		if (err != nil) {
			log.Fatalln("got error while reading request")
			return;
		}

		userInfo := &contract.User{}
		err = proto.Unmarshal(body, userInfo)

		if (err != nil) {
			log.Fatalln("got error while unmarsling")
			return;
		}

		user := model.User{}
		user.UserName = *userInfo.UserName
		user.Password = *userInfo.Password
		emailId, err := database.Login(db, &user);
		if (err != nil) {
			log.Fatalln("got error while featching user")
			res.Write([]byte("got error while featching user"))
			return
		}
		if(emailId==""){
			res.WriteHeader(http.StatusForbidden)
			return
		}
		cookieLife := time.Now().Add(-365 * 24 * time.Hour)
		cookie := http.Cookie{
			Name:"taskManagerLogin",
			Value:emailId,
			Secure:true,
			Expires:cookieLife,
		}
		http.SetCookie(res, &cookie)
		return

	}
}

func Logout(res http.ResponseWriter, req *http.Request) {
	cookieLife := time.Now().AddDate(-3, 0, 0)
	cookie := http.Cookie{
		Name:"taskManagerLogin",
		Secure:true,
		Expires:cookieLife,
	}
	http.SetCookie(res, &cookie)
	return
}