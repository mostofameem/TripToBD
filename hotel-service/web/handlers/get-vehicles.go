package handlers

import (
	"hotel-service/logger"
	"hotel-service/web/utils"
	"log/slog"
	"net/http"
	"strconv"
)

func (handlers *Handlers) GetHotel(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		slog.Error("missing id parameter in query")
		utils.SendError(w, http.StatusBadRequest, "Missing 'id' parameter in query", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error("invalid id parameter", logger.Extra(map[string]any{
			"id":  idStr,
			"err": err.Error(),
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid 'id' parameter, must be an integer", nil)
		return
	}

	vehicle, err := handlers.hotelSvc.GetHotelById(r.Context(), id)
	if err != nil {
		slog.Error("failed to get vehicle info by id", logger.Extra(map[string]any{
			"id":  id,
			"err": err.Error(),
		}))
		utils.SendError(w, http.StatusInternalServerError, "Failed to get vehicle information", nil)
		return
	}

	utils.SendData(w, vehicle)
}

func (handlers *Handlers) GetHotels(w http.ResponseWriter, r *http.Request) {
	params := utils.GetPaginationParams(r)

	vehicles, err := handlers.hotelSvc.GetHotels(r.Context(), &params)
	if err != nil {
		slog.Error("failed to get vehicle info", logger.Extra(map[string]any{
			"params": params,
			"err":    err.Error(),
		}))
		utils.SendError(w, http.StatusInternalServerError, "Failed to get vehicle information", nil)
		return
	}

	utils.SendData(w, vehicles)
}

func (h *Handlers) GetVehiclesByLocation(w http.ResponseWriter, r *http.Request) {

}
