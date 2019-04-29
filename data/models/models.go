package models

import "errors"

// Object is an object.
type Object struct {
	ID          int64        `json:"id,omitempty"`
	Name        string       `json:"name" validate:"required"`
	Description string       `json:"description" validate:"required"`
	Status      ObjectStatus `json:"status" validate:"required"`
}

// ObjectStatus is the status of an object.
type ObjectStatus string

const (
	ObjectStatusCreated = ObjectStatus("CREATED")
	ObjectStatusUpdated = ObjectStatus("UPDATED")
	ObjectStatusDeleted = ObjectStatus("DELETED")
)

func (s ObjectStatus) Validate() error {
	switch s {
	case ObjectStatusCreated:
		fallthrough
	case ObjectStatusUpdated:
		fallthrough
	case ObjectStatusDeleted:
		return nil
	default:
		return errors.New("status: invalid value")
	}
}
