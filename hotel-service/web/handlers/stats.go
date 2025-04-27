package handlers

import (
	"net/http"
	"hotel-service/metrics"
	"hotel-service/web/utils"
)

func (handlers *Handlers) Stats(w http.ResponseWriter, r *http.Request) {
	cm, err := metrics.CollectMetrics("1.0.0")
	if err != nil {
		utils.SendJson(w, http.StatusOK, map[string]any{
			"status":  false,
			"message": "Success",
		})
		return
	}
	utils.SendData(w, cm)
}
