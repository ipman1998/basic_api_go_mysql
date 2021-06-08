package main

import (
	"fmt"
	"net/http"
	"api/helper"
	"github.com/gorilla/mux"
)

func main() {
	helper.InitMessage()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/message/get", helper.GetAllMessage).Methods("GET")
	router.HandleFunc("/api/v1/message/post", helper.CreateMessage).Methods("POST")
	router.HandleFunc("/api/v1/message/find", helper.FindMessage).Methods("GET")
	router.HandleFunc("/api/v1/message/update", helper.UpdateMessage).Methods("PUT")
	router.HandleFunc("/api/v1/message/delete", helper.DeleteMessage).Methods("DELETE")
	fmt.Printf("Golang Rest API Is Running On Port: 3000")

	err := http.ListenAndServe(":3000", router)

	if err != nil {
		panic(err)
	}
}
