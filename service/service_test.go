package service_test

import (
	"database/sql"
	"net/http"
	"testing"

	"github.com/chenpengcheng/boilerplates/data/models"
	"github.com/chenpengcheng/boilerplates/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
)

//
// -- Suite
//
type Suite struct {
	suite.Suite
	dbAddr string
	db     *sql.DB
	req    *http.Request
	svc    service.Interface
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{
		dbAddr: "username:password@(localhost:3306)/product",
	})
}

func (s *Suite) SetupSuite() {
	db, err := sql.Open("mysql", s.dbAddr)
	s.NoError(err)
	s.NoError(db.Ping())

	req, err := http.NewRequest("POST", "localhost:8080", nil)
	s.NoError(err)

	svc, err := service.New(service.Config{DBAddr: s.dbAddr})
	s.NoError(err)

	s.db, s.req, s.svc = db, req, svc
}

func (s *Suite) TearDownSuite() {
	_, err := s.db.ExecContext(s.req.Context(), "DELETE from objects")
	s.NoError(err)

	s.NoError(s.db.Close())
	s.NoError(s.svc.Close())
}

//
// -- Object
//
func (s *Suite) CreateObject(object models.Object) models.Object {
	input := &service.CreateObjectInput{Object: object}
	output := &service.CreateObjectOutput{}

	s.NoError(s.svc.CreateObject(s.req, input, output))
	s.NotNil(output)
	s.NotNil(output.Object)
	s.VerifyObject(input.Object, *output.Object)

	return *output.Object
}

func (s *Suite) GetObject(ID int64) models.Object {
	input := &service.GetObjectInput{ID: ID}
	output := &service.GetObjectOutput{}

	s.NoError(s.svc.GetObject(s.req, input, output))
	s.NotNil(output)
	s.NotNil(output.Object)

	return *output.Object
}

func (s *Suite) UpdateObject(object models.Object) models.Object {
	input := &service.UpdateObjectInput{Object: object}
	output := &service.UpdateObjectOutput{}

	s.NoError(s.svc.UpdateObject(s.req, input, output))
	s.NotNil(output)
	s.NotNil(output.Object)
	s.Equal(input.Object.ID, output.Object.ID)
	s.VerifyObject(input.Object, *output.Object)

	return *output.Object
}

func (s *Suite) DeleteObject(ID int64) {
	input := &service.DeleteObjectInput{ID: ID}
	output := &service.DeleteObjectOutput{}

	s.NoError(s.svc.DeleteObject(s.req, input, output))
	s.NotNil(output)
	s.True(output.Success)
}

func (s *Suite) VerifyObject(expected, actual models.Object) {
	s.Equal(expected.Name, actual.Name)
	s.Equal(expected.Description, actual.Description)
	s.Equal(expected.Status, actual.Status)
}

//
// -- Test
//
func (s *Suite) TestCRUDObjects() {
	var createdObject models.Object

	s.Run("create an object", func() {
		object := models.Object{
			Name:        "New Object",
			Description: "This is a new object",
			Status:      models.ObjectStatusCreated,
		}

		createdObject = s.CreateObject(object)
		s.VerifyObject(object, createdObject)
	})

	s.Run("read an object", func() {
		object := s.GetObject(createdObject.ID)
		s.VerifyObject(createdObject, object)
	})

	s.Run("update an object", func() {
		object := models.Object{
			ID:          createdObject.ID,
			Name:        "Updated Object",
			Description: "This is an updated object",
			Status:      models.ObjectStatusUpdated,
		}

		updatedObject := s.UpdateObject(object)
		s.VerifyObject(object, updatedObject)
	})

	s.Run("delete an object", func() {
		s.DeleteObject(createdObject.ID)
	})
}
