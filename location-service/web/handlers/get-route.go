package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"location-service/web/utils"
// 	"strconv"
// )

// func (h *Handlers) GetRoute(w http.ResponseWriter, r *http.Request) {
// 	queryParams := r.URL.Query()

// 	locationID, err := strconv.Atoi(queryParams.Get("id"))

// 	if err != nil {
// 		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("required id"))
// 		return
// 	}

// 	locations, err := h.routeSvc.GetRoutes(r.Context(), locationID)
// 	if err != nil {
// 		utils.SendError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	utils.SendData(w, locations)

// }
