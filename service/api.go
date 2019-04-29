package service

import (
	"net/http"

	"github.com/chenpengcheng/boilerplates/data/models"
)

var (
	Name           = "Service"
	_    Interface = &service{}
)

// Config the is configuration of service.
type Config struct {
	DBAddr string `conf:"db-addr"`
}

// Interface is the interface of service.
type Interface interface {
	CreateObject(*http.Request, *CreateObjectInput, *CreateObjectOutput) error
	GetObject(*http.Request, *GetObjectInput, *GetObjectOutput) error
	UpdateObject(*http.Request, *UpdateObjectInput, *UpdateObjectOutput) error
	DeleteObject(*http.Request, *DeleteObjectInput, *DeleteObjectOutput) error
	Close() error
}

//
// -- CreateObject
//
type CreateObjectInput struct {
	Object models.Object `json:"object" validate:"required"`
}

type CreateObjectOutput struct {
	Object *models.Object `json:"object"`
}

//
// -- GetObject
//
type GetObjectInput struct {
	ID int64 `json:"id,omitempty" validate:"required"`
}

type GetObjectOutput struct {
	Object *models.Object `json:"object"`
}

//
// -- UpdateObject
//
type UpdateObjectInput struct {
	Object models.Object `json:"object" validate:"required"`
}

type UpdateObjectOutput struct {
	Object *models.Object `json:"object"`
}

//
// -- DeleteObject
//
type DeleteObjectInput struct {
	ID int64 `json:"id,omitempty" validate:"required"`
}

type DeleteObjectOutput struct {
	Success bool `json:"success"`
}
