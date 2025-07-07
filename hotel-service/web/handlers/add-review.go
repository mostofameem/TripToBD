package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"hotel-service/logger"
	"hotel-service/types"
	"hotel-service/web/utils"
)

func (h *Handlers) AddReview(w http.ResponseWriter, r *http.Request) {
	var AddReviewReq types.Reviews
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&AddReviewReq)
	if err != nil {
		slog.Error("Failed to decode request body", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", utils.ParseValidationErrors(err))
		return
	}
	err = utils.Validate(AddReviewReq)
	if err != nil {
		slog.Error("Invalid request body", logger.Extra(map[string]any{
			"body": AddReviewReq,
		}))
		utils.SendError(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	err = h.hotelSvc.AddReview(r.Context(), &AddReviewReq)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Failed to Add Review", nil)
		return
	}
	utils.SendData(w, "Review added successful")
}
