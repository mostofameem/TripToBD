package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"restaurant-service/restaurant"
	"restaurant-service/web/utils"
)

type ReqPayload struct {
	Title        string `json:"title" validate:"required"`
	Descriptions string `json:"content" validate:"required"`
	LocationInfo string `json:"location" validate:"required"`
	PictureUrl   string `json:"picture_url" validate:"required"`
}

func (handlers *Handlers) AddRestaurent(w http.ResponseWriter, r *http.Request) {
	var restaurantReq ReqPayload

	err := json.NewDecoder(r.Body).Decode(&restaurantReq)
	if err != nil {
		slog.Error("Failed to get body data")
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate(restaurantReq)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("invaild request body"))
		return
	}
	err = handlers.restSvc.AddRestaurent(r.Context(), &restaurant.Restaurant{
		Title:        restaurantReq.Title,
		LocationInfo: restaurantReq.LocationInfo,
		Descriptions: restaurantReq.Descriptions,
		PictureUrl:   restaurantReq.PictureUrl,
	})
	if err != nil {
		slog.Error(err.Error())
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, "Location addedd Successfully")
}
