package handlers

import (
	"encoding/json"
	"hotel-service/logger"
	"hotel-service/types"
	"hotel-service/web/utils"
	"log/slog"
	"net/http"
)

func (handlers *Handlers) AddHotel(w http.ResponseWriter, r *http.Request) {
	var AddVehiclesReq types.Vehicles
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&AddVehiclesReq)
	if err != nil {
		slog.Error("Failed to decode request body", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", utils.ParseValidationErrors(err))
		return
	}

	err = utils.Validate(AddVehiclesReq)
	if err != nil {
		slog.Error("Invalid request body", logger.Extra(map[string]any{
			"body": AddVehiclesReq,
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = handlers.hotelSvc.AddHotel(r.Context(), &AddVehiclesReq)
	if err != nil {
		slog.Error("Internal server error", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		utils.SendError(w, http.StatusInternalServerError, "Internal server error", nil)
		return
	}

	utils.SendData(w, "Vehicle addedd Successfully")
}
