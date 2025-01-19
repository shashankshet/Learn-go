package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shashankshet/go_api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/tracking", handlers.AddTrackingRecord)
	r.GET("/tracking/:tracking_id", handlers.GetTrackingRecord)
	r.PUT("/tracking/:tracking_id", handlers.UpdateTrackingRecord)
	r.DELETE("/tracking/:tracking_id", handlers.DeleteTrackingRecord)

	return r
}
