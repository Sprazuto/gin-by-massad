package helpers

import (
	"fmt"
	"time"

	"lke-app/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

const (
	errorLogKey   = "app:error_logs"
	maxLogEntries = 100
)

// LogError stores an error message with timestamp in Redis list
func LogError(errMsg string) error {
	redisClient := db.GetRedis()
	if redisClient == nil {
		return fmt.Errorf("redis client is not initialized")
	}

	entry := fmt.Sprintf("%s - %s", time.Now().Format(time.RFC3339), errMsg)
	err := redisClient.LPush(errorLogKey, entry).Err()
	if err != nil {
		return err
	}

	// Trim list to max length
	err = redisClient.LTrim(errorLogKey, 0, maxLogEntries-1).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetRecentErrors retrieves recent error log entries from Redis
func GetRecentErrors() ([]string, error) {
	redisClient := db.GetRedis()
	if redisClient == nil {
		return nil, fmt.Errorf("redis client is not initialized")
	}

	entries, err := redisClient.LRange(errorLogKey, 0, maxLogEntries-1).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	return entries, nil
}

// LogAPIErrorMiddleware logs API requests with non-200 status codes
func LogAPIErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		statusCode := c.Writer.Status()
		if statusCode != 200 {
			errMsg := fmt.Sprintf("Method: %s, Path: %s, Status: %d", c.Request.Method, c.Request.URL.Path, statusCode)
			_ = LogError(errMsg)
		}
	}
}
