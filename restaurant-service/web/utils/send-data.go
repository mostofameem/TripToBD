package utils

import "net/http"

func SendData(w http.ResponseWriter, data interface{}) {
	SendJSon(w, http.StatusOK, map[string]any{
		"status":  true,
		"message": "Success",
		"data":    data,
	})

}
