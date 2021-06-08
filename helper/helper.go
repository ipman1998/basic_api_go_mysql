package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"api/database"
	"api/model"
)

func InitMessage() {
	var message model.Message
	
	for i := 0; i < 10; i++ {
		message.Id = i
		message.Message = "message number" + strconv.Itoa(i)
		dal.InsertMessage(message)
	}
}

func CreateMessage(response http.ResponseWriter, request *http.Request) {
	var message model.Message
	
	json.NewDecoder(request.Body).Decode(&message)
	dal.InsertMessage(message)
}

func FindMessage(response http.ResponseWriter, request *http.Request) {
	var message model.Message
	
	json.NewDecoder(request.Body).Decode(&message)
	result := dal.GetMessage(message.Id)
	responseWithJSON(response, http.StatusOK, result)
}

func GetAllMessage(response http.ResponseWriter, request *http.Request) {
	result := dal.GetAllMessage();

	fmt.Println(strings.Trim(fmt.Sprint(result), "[]"))
	responseWithJSON(response, http.StatusOK, result)
}

func UpdateMessage(response http.ResponseWriter, request *http.Request) {
	var message model.Message

	json.NewDecoder(request.Body).Decode(&message)
	dal.UpdateMessage(message)
	data := dal.GetMessage(message.Id)
	responseWithJSON(response, http.StatusOK, data)
}

func DeleteMessage (response http.ResponseWriter, request *http.Request){
	var message model.Message

	json.NewDecoder(request.Body).Decode(&message)
	dal.DeleteMessage(message.Id)
	responseWithJSON(response, http.StatusOK, "Delete successfully")
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

// func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
// 	responseWithJSON(response, statusCode, map[string]string{
// 		"error": msg,
// 	})
// }