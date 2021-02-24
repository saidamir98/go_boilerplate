package repo

import (
	"go_boilerplate/go_boilerplate_modules/application_service"
)

// ApplicationStorageI ...
type ApplicationStorageI interface {
	Create(entity application_service.CreateApplicationModel) (res application_service.ApplicationCreatedModel, err error)
	GetList(queryParam application_service.ApplicationQueryParamModel) (res application_service.ApplicationListModel, err error)
	GetByID(id string) (res application_service.ApplicationModel, err error)
	Update(entity application_service.UpdateApplicationModel) (rowsAffected int64, err error)
	Delete(id string) (rowsAffected int64, err error)
}
