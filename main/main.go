package main

import (
	"employees-sample/controller"
	"employees-sample/repository"
	"log"
	"net/http"
)

func main() {
	err := repository.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", controller.GetEmployeeByID)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
