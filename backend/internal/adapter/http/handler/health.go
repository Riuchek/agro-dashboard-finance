package handler

import (
	"net/http"

	"github.com/Riuchek/agro-dashboard-finance/internal/adapter/http/response"
)

func Health(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
