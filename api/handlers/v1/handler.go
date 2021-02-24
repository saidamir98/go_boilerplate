package v1

import (
	"go_boilerplate/config"
	"go_boilerplate/go_boilerplate_modules/response"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/storage"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Handler ...
type Handler struct {
	cfg             config.Config
	log             logger.Logger
	storagePostgres storage.PostgresStorageI
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB) *Handler {
	return &Handler{
		cfg:             cfg,
		log:             log,
		storagePostgres: storage.NewStoragePostgres(db),
	}
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, response.SuccessModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, response.ErrorModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}

func (h *Handler) parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", h.cfg.DefaultOffset))
}

func (h *Handler) parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", h.cfg.DefaultLimit))
}
