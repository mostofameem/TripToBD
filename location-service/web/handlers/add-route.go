package handlers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log/slog"
// 	"net/http"
// 	"location-service/route"
// 	"location-service/web/utils"
// )

// type AddRouteReq struct {
// 	LocationId int            `json:"location_id" validate:"required"`
// 	Routes     [][]route.Routes `json:"routes" validate:"required"`
// }

// func (h *Handlers) AddRoute(w http.ResponseWriter, r *http.Request) {
// 	var req AddRouteReq
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		slog.Error("Failed to get body data")
// 		utils.SendError(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	err = utils.Validate(req)
// 	if err != nil {
// 		utils.SendError(w, http.StatusBadRequest, fmt.Errorf("invaild request body"))
// 		return
// 	}

// 	// get user info from grpc extract id from grpc
// 	err = h.routeSvc.AddRoutes(r.Context(), &route.RouteInfo{
// 		LocationId: req.LocationId,
// 		Route:      req.Routes,
// 	})
// 	if err != nil {
// 		utils.SendError(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	utils.SendData(w, "Thank you for your contribution")
// }
