package v1

import (
	"encoding/json"
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
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

	b, err := json.Marshal(entity)

	err = h.rmq.Push("application", "application.create", amqp.Publishing{
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent,
		ReplyTo:       "application.created", // it is used for replaying result of the event
		CorrelationId: entity.ID,
		Body:          b,
	})

	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err)
		return
	}

	h.handleSuccessResponse(c, 201, "application is being created", entity)
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
// @Success 200 {object} response.SuccessModel{data=application_service.ApplicationUpdatedModel} "Success"
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

	entity.ID = c.Param("id")

	if !util.IsValidUUID(entity.ID) {
		h.handleErrorResponse(c, 422, "validation error", "id")
		return
	}

	b, err := json.Marshal(entity)

	err = h.rmq.Push("application", "application.update", amqp.Publishing{
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent,
		CorrelationId: entity.ID,
		Body:          b,
	})

	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err)
		return
	}

	h.handleSuccessResponse(c, 200, "application is being updated", entity)
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
// @Success 200 {object} response.SuccessModel{data=application_service.DeleteApplicationModel} "Success"
// @Response 422 {object} response.ErrorModel{error=string} "Validation Error"
// @Response 400 {object} response.ErrorModel "Bad Request"
// @Failure 500 {object} response.ErrorModel "Server Error"
func (h *Handler) DeleteApplication(c *gin.Context) {
	var (
		entity application_service.DeleteApplicationModel
	)

	entity.ID = c.Param("id")

	if !util.IsValidUUID(entity.ID) {
		h.handleErrorResponse(c, 422, "validation error", "id")
		return
	}

	b, err := json.Marshal(entity)

	err = h.rmq.Push("application", "application.delete", amqp.Publishing{
		ContentType:   "application/json",
		DeliveryMode:  amqp.Persistent,
		CorrelationId: entity.ID,
		Body:          b,
	})

	if err != nil {
		h.handleErrorResponse(c, 500, "server error", err)
		return
	}

	h.handleSuccessResponse(c, 200, "application is being deleted", entity)
	return
}
