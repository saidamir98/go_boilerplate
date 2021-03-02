package api

import (
	_ "go_boilerplate/api/docs" // This should be imported for documentation
	v1 "go_boilerplate/api/handlers/v1"
	"go_boilerplate/config"
	"go_boilerplate/pkg/cors"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/pkg/pubsub"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title Go Boilerplate API
// @version 1.0
// @description This is a Go Boilerplate for medium sized projects
// @contact.name Saidamir Botirov
// @contact.email saidamir.botirov@gmail.com
// @contact.url https://www.linkedin.com/in/saidamir-botirov-a08559192
func New(cfg config.Config, log logger.Logger, db *sqlx.DB, rmq *pubsub.RMQ) (*gin.Engine, error) {
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery

	router.Use(cors.MyCORSMiddleware())

	handlerV1 := v1.New(cfg, log, db, rmq)

	router.GET("/ping", handlerV1.Ping)
	router.GET("/config", handlerV1.GetConfig)

	rV1 := router.Group("/v1")
	{
		endpointsV1(rV1, handlerV1)
	}

	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router, nil
}
