package handlers

import (
	"bytes"
	"encoding/json"
	service_mock "github.com/code7unner/rest-api-test-task/users/internal/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AuthSuite struct {
	suite.Suite
	mockServiceCtl *gomock.Controller
	mockService    *service_mock.MockService
}

func (s *AuthSuite) SetupTest() {
	s.mockServiceCtl = gomock.NewController(s.T())
	s.mockService = service_mock.NewMockService(s.mockServiceCtl)
}

func (s *AuthSuite) TearDownTest() {
	s.mockServiceCtl.Finish()
}

func (s *AuthSuite) buildRequest(data []byte) *http.Request {
	req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")

	return req
}

func (s *AuthSuite) TestRegister() {
	reqStruct := RegisterRequest{
		Username: "TestUsername",
		Password: "TestPass",
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		s.T().Fail()
	}

	req := s.buildRequest(body)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/register")

	h := NewAuthHandler(s.mockService)
	s.mockService.
		EXPECT().
		Register(reqStruct.Username, reqStruct.Password).
		Return(nil, nil)
	s.Require().NoError(h.Register(c))
	s.Require().Equal(http.StatusCreated, rec.Code)
}

func (s *AuthSuite) TestLogin() {
	reqStruct := RegisterRequest{
		Username: "TestUsername",
		Password: "TestPass",
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		s.T().Fail()
	}

	req := s.buildRequest(body)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/login")

	h := NewAuthHandler(s.mockService)
	s.mockService.
		EXPECT().
		Login(reqStruct.Username, reqStruct.Password).
		Return("", nil)

	s.Require().NoError(h.Login(c))
	s.Require().Equal(http.StatusOK, rec.Code)
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}
