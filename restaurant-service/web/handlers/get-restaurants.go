package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"restaurant-service/logger"
	"restaurant-service/web/utils"
	"strconv"
	"strings"
)

func (handlers *Handlers) GetRestaurant(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		slog.Error("Restaurant ID is required", logger.Extra(map[string]any{
			"patPart": pathParts,
		}))
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("no id found"))
		return
	}

	restaurantID, err := strconv.Atoi(pathParts[2])
	if err != nil {
		slog.Error("Invalid restaurant ID format")
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("invalid id"))
		return
	}

	restaurant, err := handlers.restSvc.GetRestaurant(r.Context(), restaurantID)
	if err != nil {
		slog.Error("Error retrieving restaurant data", "error", err)
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, restaurant)
}

func (handlers *Handlers) GetRestaurants(w http.ResponseWriter, r *http.Request) {
	paginationParams := utils.GetPaginationParams(r, "ASC", "rating")

	restaurents, err := handlers.restSvc.GetRestaurants(r.Context(), paginationParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, restaurents)
}

func (h *Handlers) GetRestaurantsByLocation(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 4 || pathParts[3] == "" {
		slog.Error("id required")
		utils.SendError(w, http.StatusBadRequest, nil)
		return
	}

	locationId, err := strconv.Atoi(pathParts[3])
	if err != nil {
		slog.Error("Invalid id")
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	paginationParams := utils.GetPaginationParams(r, "ASC", "rating")

	restaurants, err := h.restSvc.GetRestaurantsByLocation(r.Context(), locationId, paginationParams)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, restaurants)
}
