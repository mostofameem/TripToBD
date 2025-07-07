package handlers

import (
	"fmt"
	"location-service/web/utils"
	"net/http"
)

func (h *Handlers) GetRoute(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	locationID := queryParams.Get("id")

	if locationID == "" {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("required id"))
		return
	}

	locations, err := h.locSvc.GetRoutes(r.Context(), locationID)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, locations)

}
