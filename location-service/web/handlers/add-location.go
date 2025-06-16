package handlers

import (
	"encoding/json"
	"fmt"
	"location-service/location"
	"location-service/web/utils"
	"log/slog"
	"net/http"
	"time"
)

type LocationReq struct {
	Title      string `json:"title"                validate:"required"`
	Details    string `json:"details"              validate:"required"`
	BestTime   string `json:"bestTime"             validate:"required"`
	PictureUrl string `json:"pictureUrl"           validate:"required"`
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

	err = handlers.locSvc.AddLocation(r.Context(), &location.AddLocationReq{
		Title:        locationReq.Title,
		Descriptions: locationReq.Details,
		BestTime:     locationReq.BestTime,
		PictureUrl:   locationReq.PictureUrl,
		Rating:       0,
		Voted:        0,
		CreatedBy:    1,
		CreatedAt:    time.Now(),
	})
	if err != nil {
		slog.Error(err.Error())
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, "Location addedd Successfully")
}
