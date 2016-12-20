package main

import (
	"taskManagerUserService/database"
	"log"
	"net/http"
	"taskManagerUserService/router"
)

func main() {
	db, err := database.OpenDatabase()
	if (err != nil) {
		log.Fatal(err.Error())
	}
	defer db.Close()
	router.HandleRequest(db);
	http.ListenAndServe(":5000", nil)
}
