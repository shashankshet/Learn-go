package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shashankshet/go_api/db"
)

type TrackingRequest struct {
	InputPath  string `json:"input_path" binding:"required"`
	OutputPath string `json:"output_path" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

func AddTrackingRecord(c *gin.Context) {
	var request TrackingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trackingID := uuid.New().String()
	timestamp := time.Now()

	_, err := db.DB.Exec(`
		INSERT INTO tracking (tracking_id, "timestamp", input_path, output_path, status)
		VALUES ($1, $2, $3, $4, $5)`,
		trackingID, timestamp, request.InputPath, request.OutputPath, request.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tracking_id": trackingID,
		"timestamp":   timestamp,
	})
}

func GetTrackingRecord(c *gin.Context) {
	trackingID := c.Param("tracking_id")

	var (
		inputPath  string
		outputPath string
		status     string
		timestamp  time.Time
	)

	err := db.DB.QueryRow(`
		SELECT "timestamp", input_path, output_path, status
		FROM tracking
		WHERE tracking_id = $1`, trackingID).Scan(&timestamp, &inputPath, &outputPath, &status)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tracking_id": trackingID,
		"timestamp":   timestamp,
		"input_path":  inputPath,
		"output_path": outputPath,
		"status":      status,
	})
}

func UpdateTrackingRecord(c *gin.Context) {
	trackingID := c.Param("tracking_id")

	var request TrackingRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.DB.Exec(`
		UPDATE tracking
		SET input_path = $1, output_path = $2, status = $3
		WHERE tracking_id = $4`,
		request.InputPath, request.OutputPath, request.Status, trackingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}

func DeleteTrackingRecord(c *gin.Context) {
	trackingID := c.Param("tracking_id")

	_, err := db.DB.Exec(`
		DELETE FROM tracking
		WHERE tracking_id = $1`, trackingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
