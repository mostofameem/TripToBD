package handlers

import (
	"fmt"
	"net/http"
	"post-service/web/utils"
	"strconv"
)

func (handlers *Handlers) GetLocation(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	locationId := queryParams.Get("id")
	id, err := strconv.Atoi(locationId)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("required id"))
		return
	}

	locations, err := handlers.locSvc.GetLocation(r.Context(), id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, locations)
}

func (handlers *Handlers) GetLocations(w http.ResponseWriter, r *http.Request) {
	paginationParams := utils.GetPaginationParams(r, "ASC", "rating")
	locations, err := handlers.locSvc.GetLocations(r.Context(), paginationParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}
	utils.SendData(w, locations)
}
