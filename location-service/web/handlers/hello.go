package handlers

import (
	"net/http"
	"post-service/web/utils"
)

func (handlers *Handlers) Hello(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, "It's Running !!!")
}
