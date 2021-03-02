package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "go_boilerplate/api/handlers/v1"
)

func endpoints(r *gin.Engine, h *v1.Handler) {
	r.GET("/ping", h.Ping)
	r.GET("/config", h.GetConfig)

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	rv1 := r.Group("/v1")
	{
		rv1.POST("/application", h.CreateApplication)
		rv1.GET("/application", h.GetApplicationList)
		rv1.GET("/application/:id", h.GetApplicationByID)
		rv1.PUT("/application/:id", h.UpdateApplication)
		rv1.DELETE("/application/:id", h.DeleteApplication)
	}
}