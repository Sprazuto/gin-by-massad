package controllers

import (
	"net/http"
	"os"

	"lke-app/helpers"

	"github.com/gin-gonic/gin"
)

// AppController handles application-level endpoints
// @title App Controller
// @version 1.0
// @description Application version and system information endpoints
type AppController struct{}

// Version returns the current version of the app from environment variable
// @Summary Get application version
// @Description Get current application version and patch information
// @Tags Application
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{} "Application version information"
// @Router /v1/version [get]
func (a *AppController) Version(c *gin.Context) {
	version := os.Getenv("API_VERSION")
	if version == "" {
		version = "unknown"
	}
	c.JSON(http.StatusOK, gin.H{
		"version": version,
		"patch":   "1.7.15",
	})
}

// ErrorLogs returns recent error logs from Redis for debugging
// @Summary Get recent error logs
// @Description Retrieve recent application error logs for debugging
// @Tags Application
// @Produce json
// @Security Bearer
// @Success 200 {object} map[string]interface{} "Recent error logs"
// @Failure 500 {object} map[string]interface{} "Failed to get error logs"
// @Router /v1/error-logs [get]
func (a *AppController) ErrorLogs(c *gin.Context) {
	entries, err := helpers.GetRecentErrors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get error logs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error_logs": entries})
}
