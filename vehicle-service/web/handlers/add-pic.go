package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"vehicles/logger"
	"vehicles/types"
	"vehicles/web/utils"
)

func (h *Handlers) AddPic(w http.ResponseWriter, r *http.Request) {
	var AddPicReq types.Pictures
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&AddPicReq)
	if err != nil {
		slog.Error("Failed to decode request body", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", utils.ParseValidationErrors(err))
		return
	}

	err = utils.Validate(AddPicReq)
	if err != nil {
		slog.Error("Invalid request body", logger.Extra(map[string]any{
			"body": AddPicReq,
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.vehicleSvc.AddPics(r.Context(), &AddPicReq)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Failed to Add Pics", nil)
		return
	}
	utils.SendData(w, "Pics added successful")
}

