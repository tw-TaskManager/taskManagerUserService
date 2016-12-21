package router

import "database/sql"
import "github.com/gorilla/mux"
import (
	userHandler "taskManagerUserService/handler"
	"net/http"
)

func HandleRequest(db *sql.DB) {
	handler := mux.NewRouter()
	handler.HandleFunc("/task/createUser", userHandler.CreateUserTask(db)).Methods("POST")
	handler.HandleFunc("/task/login", userHandler.LoginUser(db)).Methods("POST")
	handler.HandleFunc("/task/logout", userHandler.Logout).Methods("POST")
	handler.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	http.Handle("/", handler)
}
