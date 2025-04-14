package handlers

import (
	"net/http"
	"vehicles/web/utils"
)

func (h *Handlers) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	vehicleId := utils.ParseId(r)

	if vehicleId == 0 {
		utils.SendError(w, http.StatusBadRequest, "id required", nil)
		return
	}

	reviews, err := h.vehicleSvc.GetReviewsById(r.Context(), vehicleId)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Failed to get Reviews", nil)
		return
	}

	utils.SendData(w, reviews)
}
