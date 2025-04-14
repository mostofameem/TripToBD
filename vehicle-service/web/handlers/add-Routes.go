package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"vehicles/logger"
	"vehicles/types"
	"vehicles/web/utils"
)

func (h *Handlers) AddRoutes(w http.ResponseWriter, r *http.Request) {
	var AddRoutesReq types.Routes
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&AddRoutesReq)
	if err != nil {
		slog.Error("Failed to decode request body", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", utils.ParseValidationErrors(err))
		return
	}

	// vehicleId, err := middlewares.GetVehicleId(r)
	// if err != nil {
	// 	slog.Error("Failed to get Id from header", logger.Extra(map[string]any{
	// 		"error": err.Error(),
	// 	}))
	// 	utils.SendError(w, http.StatusBadRequest, "Unauthorized", nil)
	// 	return
	// }
	// AddRoutesReq.VehicleId = vehicleId

	err = utils.Validate(AddRoutesReq)
	if err != nil {
		slog.Error("Invalid request body", logger.Extra(map[string]any{
			"body": AddRoutesReq,
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.vehicleSvc.AddRoute(r.Context(), &AddRoutesReq)
	if err != nil {
		slog.Error("Failed to insert route", logger.Extra(map[string]any{
			"body": AddRoutesReq,
		}))
		utils.SendError(w, http.StatusBadRequest, "Failed to insert route", nil)
		return
	}
	utils.SendData(w, "Routes added successful")
}
