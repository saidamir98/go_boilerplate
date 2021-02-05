package repo

import "go_boilerplate/api/models"

// ApplicationStorageI ...
type ApplicationStorageI interface {
	Create(entity models.CreateApplication) (res models.ApplicationCreated, err error)
	GetList(queryParam models.ApplicationQueryParam) (res models.ApplicationList, err error)
	GetByID(id string) (res models.Application, err error)
	Update(entity models.UpdateApplication) (rowsAffected int64, err error)
	Delete(id string) (rowsAffected int64, err error)
}
