package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"restaurant-service/restaurant"
	"restaurant-service/web/utils"
)

type addReviewReq struct {
	RestaurantId int                `json:"restaurant_id"`
	Comment      restaurant.Comment `json:"comment" validate:"required"`
}

func (h *Handlers) AddReview(w http.ResponseWriter, r *http.Request) {
	var req addReviewReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("Failed to get body data")
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate(req)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("invaild request body"))
		return
	}

	err = h.restSvc.AddReviews(r.Context(), req.RestaurantId, req.Comment)
	if err != nil {
		slog.Error("Failed to add branch to the location")
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, "Review added successful")
}
