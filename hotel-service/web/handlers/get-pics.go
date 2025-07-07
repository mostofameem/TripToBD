package handlers

import (
	"net/http"
	"hotel-service/web/utils"
)

func (h *Handlers) GetPic(w http.ResponseWriter, r *http.Request) {
	vehicleId := utils.ParseId(r)

	if vehicleId == 0 {
		utils.SendError(w, http.StatusBadRequest, "id required", nil)
		return
	}

	pics, err := h.hotelSvc.GetPicsById(r.Context(), vehicleId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to fetth Pics", nil)
		return
	}
	utils.SendData(w, pics)
}
