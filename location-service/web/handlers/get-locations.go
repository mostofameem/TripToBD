package handlers

import (
	"fmt"
	"location-service/web/utils"
	"net/http"
)

func (handlers *Handlers) GetLocation(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	locationId := queryParams.Get("id")

	if locationId == "" {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("required id"))
		return
	}

	locations, err := handlers.locSvc.GetLocation(r.Context(), locationId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, locations)
}

func (handlers *Handlers) GetLocations(w http.ResponseWriter, r *http.Request) {
	paginationParams := utils.GetPaginationParams(r, "DESC", "rating")
	locations, err := handlers.locSvc.GetLocationPage(r.Context(), paginationParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, locations)
}
