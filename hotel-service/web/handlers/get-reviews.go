package handlers

import (
	"net/http"
	"hotel-service/web/utils"
)

func (h *Handlers) GetReview(w http.ResponseWriter, r *http.Request) {
	vehicleId := utils.ParseId(r)

	if vehicleId == 0 {
		utils.SendError(w, http.StatusBadRequest, "id required", nil)
		return
	}

	reviews, err := h.hotelSvc.GetReviewsById(r.Context(), vehicleId)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Failed to get Reviews", nil)
		return
	}

	utils.SendData(w, reviews)
}
