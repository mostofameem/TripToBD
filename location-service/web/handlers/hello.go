package handlers

import (
	"location-service/web/utils"
	"net/http"
)

func (handlers *Handlers) Hello(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, "It's Running !!!")
}
