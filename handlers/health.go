package handlers

import (
	"go-ms-adapter/utils"
	"net/http"
)

// HealthCheck :Check for health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "Healthy!")
}
