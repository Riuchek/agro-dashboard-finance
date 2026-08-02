package handler

import (
	"net/http"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/http/response"
	"github.com/Riuchek/agro-dashboard-finance/internal/app/overview"
)

type Overview struct {
	service *overview.Service
}

func NewOverview(service *overview.Service) *Overview {
	return &Overview{service: service}
}

func (h *Overview) Get(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, h.service.Get(r.Context()))
}
