package handlers

import (
	"encoding/json"
	"fmt"
	"location-service/location"
	"location-service/web/utils"
	"log/slog"
	"net/http"
)

type AddRouteReq struct {
	LocationId string         `json:"locationId"  validate:"required"`
	Routes     location.Route `json:"routes"      validate:"required"`
}

func (h *Handlers) AddRoute(w http.ResponseWriter, r *http.Request) {
	var req AddRouteReq
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

	err = h.locSvc.AddRoutes(r.Context(), req.LocationId, &req.Routes)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, "Thank you for your contribution")
}
