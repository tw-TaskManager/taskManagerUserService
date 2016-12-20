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
	http.Handle("/",handler)
}
