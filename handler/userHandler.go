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
	"fmt"
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
			log.Fatalln("got error while creating user")
			res.Write([]byte("got error while creating user"))
			return
		}

		res.Write([]byte("user created"))

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
		fmt.Println("email id",emailId)
		if (err != nil) {
			log.Fatalln("got error while featching user")
			res.Write([]byte("got error while featching user"))
			return
		}

		response := &contract.Response{}
		response.Response = []byte(emailId)
		emailId_to_send, err := proto.Marshal(response)

		if (err != nil) {
			log.Fatalln("got error while marsling user")
			res.Write([]byte("got error while marsling user"))
			return
		}
		res.Write([]byte(emailId_to_send))

	}
}