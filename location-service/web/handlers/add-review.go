package handlers

// import (
// 	"encoding/json"
// 	"log/slog"
// 	"net/http"
// 	"location-service/location"
// 	"location-service/web/utils"
// )

// func (handlers *Handlers) AddReview(w http.ResponseWriter, r *http.Request) {
// 	var req location.Comment
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		slog.Error("Failed to get body data")
// 		utils.SendError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	err = handlers.locSvc.AddReviews(r.Context(), &req)
// 	if err != nil {
// 		utils.SendError(w, http.StatusInternalServerError, err)
// 	}
	
// 	utils.SendData(w, "Thank you for your contribution!!")
// }
