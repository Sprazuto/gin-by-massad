package controllers

import (
	"net/http"
	"os"

	"aset-app/helpers"

	"github.com/gin-gonic/gin"
)

type AppController struct{}

// Version returns the current version of the app from environment variable
func (a *AppController) Version(c *gin.Context) {
	version := os.Getenv("API_VERSION")
	if version == "" {
		version = "unknown"
	}
	c.JSON(http.StatusOK, gin.H{
		"version": version,
		"patch":   "1.0.3.5",
	})
}

// ErrorLogs returns recent error logs from Redis for debugging
func (a *AppController) ErrorLogs(c *gin.Context) {
	entries, err := helpers.GetRecentErrors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get error logs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"error_logs": entries})
}
