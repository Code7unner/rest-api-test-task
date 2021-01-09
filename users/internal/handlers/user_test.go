package handlers

import (
	"bytes"
	service_mock "github.com/code7unner/rest-api-test-task/users/internal/service/mock"
	"github.com/code7unner/rest-api-test-task/users/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserSuite struct {
	suite.Suite
	mockServiceCtl *gomock.Controller
	mockService    *service_mock.MockService
	id             string
}

func (s *UserSuite) SetupTest() {
	s.id = "10"
	s.mockServiceCtl = gomock.NewController(s.T())
	s.mockService = service_mock.NewMockService(s.mockServiceCtl)
}

func (s *UserSuite) TearDownTest() {
	s.mockServiceCtl.Finish()
}

func (s *UserSuite) buildRequest() *http.Request {
	req := httptest.NewRequest(echo.GET, "/", bytes.NewBuffer(nil))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer wqrqewrqwerqwer")

	q := req.URL.Query()
	q.Add("id", s.id)
	req.URL.RawQuery = q.Encode()

	return req
}

func (s *UserSuite) TestGetUserByID() {
	req := s.buildRequest()
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues(s.id)

	h := NewUserHandler(s.mockService)
	user := new(models.Users)
	s.mockService.EXPECT().GetUser(10).Return(user, nil)

	s.Require().NoError(h.GetUser(c))
	s.Require().Equal(http.StatusOK, rec.Code)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
