package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"post-service/location"
	"post-service/web/utils"
)

type LocationReq struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content"`
	BestTime   string `json:"best_time" validate:"required"`
	PictureUrl string `json:"picture_url" validate:"required"`
}

func (handlers *Handlers) AddLocation(w http.ResponseWriter, r *http.Request) {
	var locationReq LocationReq

	err := json.NewDecoder(r.Body).Decode(&locationReq)
	if err != nil {
		slog.Error("Failed to get body data")
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}

	err = utils.Validate(locationReq)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("invaild request body"))
		return
	}
	err = handlers.locSvc.AddLocation(r.Context(), &location.Location{
		Title:      locationReq.Title,
		Content:    locationReq.Content,
		BestTime:   locationReq.BestTime,
		PictureUrl: locationReq.PictureUrl,
	})
	if err != nil {
		slog.Error(err.Error())
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, "Location addedd Successfully")
}
