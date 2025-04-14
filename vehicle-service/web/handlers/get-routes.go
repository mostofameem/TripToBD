package handlers

import (
	"log/slog"
	"net/http"
	"vehicles/logger"
	"vehicles/types"
	"vehicles/web/utils"
)

func (h *Handlers) GetRoutesByVehicle(w http.ResponseWriter, r *http.Request) {
	params := types.GetRoutesByVehicleParams{
		Vehicle_id: utils.ParseId(r),
		Limit:      utils.ParseLimit(r),
		Page:       utils.ParsePage(r),
	}

	err := utils.Validate(params)
	if err != nil {
		slog.Error("Invalid query params", logger.Extra(map[string]any{
			"params": params,
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid query params", nil)
		return
	}

	routes, err := h.vehicleSvc.GetRouteByVehicleId(r.Context(), &params)
	if err != nil {
		slog.Error("failed to get vehicle info", logger.Extra(map[string]any{
			"routes": routes,
			"err":    err.Error(),
		}))
		utils.SendError(w, http.StatusInternalServerError, "Failed to get routes", nil)
		return
	}

	utils.SendData(w, routes)

}

func (h *Handlers) GetRoutesByLocation(w http.ResponseWriter, r *http.Request) {
	req := types.GetGetRoutesByLocationParams{
		Dest:  utils.ParseDest(r),
		Limit: utils.ParseLimit(r),
		Page:  utils.ParsePage(r),
	}

	if len(req.Dest) == 0 {
		utils.SendError(w, http.StatusBadRequest, "dest field is required", nil)
		return
	}

	routes, err := h.vehicleSvc.GetRoutesByLocation(r.Context(), &req)
	if err != nil {
		slog.Error("failed to get vehicle info", logger.Extra(map[string]any{
			"routes": routes,
			"err":    err.Error(),
		}))
		utils.SendError(w, http.StatusInternalServerError, "Failed to get routes", nil)
		return
	}

	utils.SendData(w, routes)

}
