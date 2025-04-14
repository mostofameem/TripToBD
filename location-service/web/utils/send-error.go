package utils

import "net/http"

func SendError(w http.ResponseWriter, status int, err error){
	SendJSon(w,status ,map[string]any{
		"status":false,
		"message": err.Error(),
	})
}