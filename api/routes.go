package api

import (
	v1 "go_boilerplate/api/handlers/v1"

	"github.com/gin-gonic/gin"
)

func endpointsV1(r *gin.RouterGroup, h *v1.Handler) {
	r.POST("/application", h.CreateApplication)
	r.GET("/application", h.GetApplicationList)
	r.GET("/application/:id", h.GetApplicationByID)
	r.PUT("/application/:id", h.UpdateApplication)
	r.DELETE("/application/:id", h.DeleteApplication)
}
