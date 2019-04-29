package service

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/chenpengcheng/boilerplates/data/models"
	_ "github.com/go-sql-driver/mysql"
	validator "gopkg.in/go-playground/validator.v9"
)

// service implements Interface.
type service struct {
	db *sql.DB
}

// New creates a new service.
func New(config Config) (Interface, error) {
	db, err := sql.Open("mysql", config.DBAddr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &service{db: db}, nil
}

// CreateObject creates an object.
func (s *service) CreateObject(r *http.Request, input *CreateObjectInput, output *CreateObjectOutput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	result, err := s.db.ExecContext(r.Context(),
		"INSERT INTO objects (name, description, status) VALUES(?,?,?)",
		input.Object.Name, input.Object.Description, input.Object.Status,
	)
	if err != nil {
		return err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	var getOutput GetObjectOutput
	if err := s.GetObject(r, &GetObjectInput{ID: ID}, &getOutput); err != nil {
		return err
	}
	output.Object = getOutput.Object

	return nil
}

// Validate validates CreateObjectInput.
func (i CreateObjectInput) Validate() error {
	if err := validator.New().Struct(i); err != nil {
		return err
	}
	return i.Object.Status.Validate()
}

// GetObject creates an object.
func (s *service) GetObject(r *http.Request, input *GetObjectInput, output *GetObjectOutput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	var object models.Object

	row := s.db.QueryRowContext(r.Context(),
		"SELECT id, name, description, status FROM objects WHERE id=?",
		input.ID)
	if err := row.Scan(
		&object.ID,
		&object.Name,
		&object.Description,
		&object.Status,
	); err != nil {
		return err
	}

	output.Object = &object
	return nil
}

// Validate validates GetObjectInput.
func (i GetObjectInput) Validate() error {
	return validator.New().Struct(i)
}

// UpdateObject updates an object.
func (s *service) UpdateObject(r *http.Request, input *UpdateObjectInput, output *UpdateObjectOutput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	_, err := s.db.ExecContext(r.Context(),
		"UPDATE objects SET name=?, description=?, status=? WHERE id=?",
		input.Object.Name, input.Object.Description, models.ObjectStatusUpdated,
		input.Object.ID,
	)
	if err != nil {
		return err
	}

	var getOutput GetObjectOutput
	if err := s.GetObject(r, &GetObjectInput{ID: input.Object.ID}, &getOutput); err != nil {
		return err
	}
	output.Object = getOutput.Object

	return nil
}

// Validate validates UpdateObjectInput.
func (i UpdateObjectInput) Validate() error {
	if err := validator.New().Struct(i); err != nil {
		return err
	}
	if i.Object.ID < 0 {
		return errors.New("update object: invalid ID")
	}

	return i.Object.Status.Validate()
}

// DeleteObject deletes an object.
func (s *service) DeleteObject(r *http.Request, input *DeleteObjectInput, output *DeleteObjectOutput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	result, err := s.db.ExecContext(r.Context(),
		"DELETE FROM objects WHERE id=?",
		input.ID,
	)
	if err != nil {
		return err
	}

	n, err := result.RowsAffected()
	if err != nil {
		return err
	}

	output.Success = (n >= 1)
	return nil
}

// Validate validates DeleteObjectInput.
func (i DeleteObjectInput) Validate() error {
	return validator.New().Struct(i)
}

// Close closes the DB connections openned by service.
func (s *service) Close() error {
	return s.db.Close()
}
