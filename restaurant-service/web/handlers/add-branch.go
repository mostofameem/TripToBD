package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"restaurant-service/web/utils"
)

type BranchPayload struct {
	LocationId   int `json:"location_id" validate:"required"`
	RestaurantId int `json:"restaurant_id" validate:"required"`
}

func (h *Handlers) AddBranch(w http.ResponseWriter, r *http.Request) {
	var branchInfo BranchPayload
	err := json.NewDecoder(r.Body).Decode(&branchInfo)
	if err != nil {
		slog.Error("Failed to get body data")
		utils.SendError(w, http.StatusBadRequest, err)
		return
	}
	err = utils.Validate(branchInfo)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("invaild request body"))
		return
	}
	err = h.restSvc.AddBranchs(r.Context(), branchInfo.LocationId, branchInfo.RestaurantId)
	if err != nil {
		slog.Error("Failed to add branch to the location")
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, "Branch added successful")
}
