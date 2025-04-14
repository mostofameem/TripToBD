package handlers

import (
	"database/sql"
	"net/http"
	"restaurant-service/web/utils"
)

type SearchReq struct {
	SearchTitle    string `json:"title"`
	SearchLocation string `json:"location"`
}

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	req := SearchReq{
		SearchTitle:    r.URL.Query().Get("title"),
		SearchLocation: r.URL.Query().Get("location"),
	}

	restaurants, err := h.restSvc.Search(r.Context(), req.SearchTitle, req.SearchLocation)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.SendData(w, nil)
			return
		}
		utils.SendError(w, http.StatusInternalServerError, err)
		return
	}

	utils.SendData(w, restaurants)
}
