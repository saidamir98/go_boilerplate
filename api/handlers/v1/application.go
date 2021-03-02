package v1

import (
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/pkg/logger"
	"go_boilerplate/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateApplication godoc
// @ID create-application
// @Router /v1/application [POST]
// @Tags application
// @Summary creates an application
// @Description creates an application
// @Accept json
// @Param application body application_service.CreateApplicationModel true "application body"
// @Produce json
// @Success 201 {object} response.SuccessModel{data=application_service.ApplicationCreatedModel} "Success"
// @Response 422 {object} response.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} response.ErrorModel "Bad Request"
// @Failure 500 {object} response.ErrorModel "Server Error"
func (h *Handler) CreateApplication(c *gin.Context) {
	var (
		entity application_service.CreateApplicationModel
	)

	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err)
		return
	}

	if entity.Body == "" {
		h.handleErrorResponse(c, 422, "validation error", "body")
		return
	}

	uuid, err := uuid.NewRandom()

	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err)
		return
	}

	entity.ID = uuid.String()

	res, err := h.storagePostgres.Application().Create(entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err)
		return
	}

	h.log.Info("application has been created", logger.Any("body", entity), logger.Any("result", res))

	h.handleSuccessResponse(c, 201, "application has been created", res)
	return
}

// GetApplicationList godoc
// @ID get-application-list
// @Router /v1/application [GET]
// @Tags application
// @Summary gets application list
// @Description gets application list
// @Accept json
// @Param find query application_service.ApplicationQueryParamModel false "filters"
// @Produce json
// @Success 200 {object} response.SuccessModel{data=application_service.ApplicationListModel} "Success"
// @Response 400 {object} response.ErrorModel "Bad Request"
// @Failure 500 {object} response.ErrorModel "Server Error"
func (h *Handler) GetApplicationList(c *gin.Context) {
	var (
		queryParam application_service.ApplicationQueryParamModel
		err        error
	)

	queryParam.Search = c.DefaultQuery("search", "")
	queryParam.Order = c.DefaultQuery("order", "")
	queryParam.Arrangement = c.DefaultQuery("arrangement", "")

	queryParam.Offset, err = h.parseOffsetQueryParam(c)
	if err != nil {
		h.handleErrorResponse(c, 400, "wrong offset input", err)
		return
	}
	queryParam.Limit, err = h.parseLimitQueryParam(c)
	if err != nil {
		h.handleErrorResponse(c, 400, "wrong limit input", err)
		return
	}

	res, err := h.storagePostgres.Application().GetList(queryParam)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err)
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

// GetApplicationByID godoc
// @ID get-application-by-id
// @Router /v1/application/{id} [GET]
// @Tags application
// @Summary gets an application by its id
// @Description gets an application by its id
// @Accept json
// @Param id path string true "application id"
// @Produce json
// @Success 200 {object} response.SuccessModel{data=application_service.ApplicationModel} "Success"
// @Response 422 {object} response.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} response.ErrorModel "Bad Request"
// @Failure 500 {object} response.ErrorModel "Server Error"
func (h *Handler) GetApplicationByID(c *gin.Context) {
	id := c.Param("id")

	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, 422, "validation error", "id")
		return
	}

	res, err := h.storagePostgres.Application().GetByID(id)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err)
		return
	}

	h.handleSuccessResponse(c, 200, "ok", res)
	return
}

// UpdateApplication godoc
// @ID update-application
// @Router /v1/application/{id} [PUT]
// @Tags application
// @Summary gets an application by its id
// @Description gets an application by its id
// @Accept json
// @Param id path string true "application id"
// @Param application body application_service.UpdateApplicationModel true "application body"
// @Produce json
// @Success 200 {object} response.SuccessModel{data=int64} "Success"
// @Response 422 {object} response.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} response.ErrorModel "Bad Request"
// @Failure 500 {object} response.ErrorModel "Server Error"
func (h *Handler) UpdateApplication(c *gin.Context) {
	var (
		entity application_service.UpdateApplicationModel
	)

	err := c.ShouldBindJSON(&entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err)
		return
	}

	if entity.Body == "" {
		h.handleErrorResponse(c, 422, "validation error", "body")
		return
	}

	id := c.Param("id")

	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, 422, "validation error", "id")
		return
	}

	entity.ID = id

	rowsAffected, err := h.storagePostgres.Application().Update(entity)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err)
		return
	}

	h.handleSuccessResponse(c, 200, "application has been updated", rowsAffected)
	return
}

// DeleteApplication godoc
// @ID delete-application
// @Router /v1/application/{id} [DELETE]
// @Tags application
// @Summary deletes an application by its id
// @Description deletes an application by its id
// @Accept json
// @Param id path string true "application id"
// @Produce json
// @Success 200 {object} response.SuccessModel{data=int64} "Success"
// @Response 422 {object} response.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} response.ErrorModel "Bad Request"
// @Failure 500 {object} response.ErrorModel "Server Error"
func (h *Handler) DeleteApplication(c *gin.Context) {
	id := c.Param("id")

	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, 422, "validation error", "id")
		return
	}

	rowsAffected, err := h.storagePostgres.Application().Delete(id)

	if err != nil {
		h.handleErrorResponse(c, 400, "bad request", err)
		return
	}

	if rowsAffected == 1 {
		h.log.Info("application has been deleted", logger.String("id", id))
	}

	h.handleSuccessResponse(c, 200, "application has been deleted", rowsAffected)
	return
}
