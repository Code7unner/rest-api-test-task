package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/code7unner/rest-api-test-task/todos/internal/service"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	service_mock "github.com/code7unner/rest-api-test-task/todos/internal/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type TodosSuite struct {
	suite.Suite
	mockServiceCtl *gomock.Controller
	mockService    *service_mock.MockService
	userID         int
	todoID         int
}

func (s *TodosSuite) SetupTest() {
	s.userID = 111
	s.todoID = 1
	s.mockServiceCtl = gomock.NewController(s.T())
	s.mockService = service_mock.NewMockService(s.mockServiceCtl)
}

func (s *TodosSuite) TearDownTest() {
	s.mockServiceCtl.Finish()
}

func (s *TodosSuite) buildPostRequest(data []byte) *http.Request {
	req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")

	return req
}

func (s *TodosSuite) buildPatchRequest(data []byte) *http.Request {
	req := httptest.NewRequest(echo.PATCH, "/", bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")

	return req
}

func (s *TodosSuite) buildDeleteRequest(data []byte) *http.Request {
	req := httptest.NewRequest(echo.DELETE, "/", bytes.NewBuffer(data))
	req.Header.Set("Content-type", "application/json")

	return req
}

func (s *TodosSuite) buildGetRequest() *http.Request {
	req := httptest.NewRequest(echo.GET, "/", bytes.NewBuffer(nil))
	req.Header.Set("Content-type", "application/json")

	return req
}

func (s *TodosSuite) addToken(c echo.Context) echo.Context {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = float64(s.userID)
	c.Set("user", token)

	return c
}

func (s *TodosSuite) parseTime(value string) time.Time {
	if value != "" {
		t, _ := time.Parse(service.DateTimeLayout, value)
		return t
	}

	return time.Time{}
}

func (s *TodosSuite) TestGetAllCurrentTodos() {
	reqStruct := CurrentTodoRequest{
		Time: "2010-09-22 12:42:31",
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		s.T().Fail()
	}

	req := s.buildPostRequest(body)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todo/current")

	c = s.addToken(c)

	h := NewTodosHandler(s.mockService)
	s.mockService.
		EXPECT().
		GetAllCurrentTodos(s.userID, s.parseTime(reqStruct.Time)).
		Return(nil, nil)
	s.Require().NoError(h.GetAllCurrentTodos(c))
	s.Require().Equal(http.StatusOK, rec.Code)
}

func (s *TodosSuite) TestGetAllTodos() {
	req := s.buildGetRequest()
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todo/all")

	c = s.addToken(c)

	h := NewTodosHandler(s.mockService)
	s.mockService.
		EXPECT().
		GetAllTodos(s.userID).
		Return(nil, nil)
	s.Require().NoError(h.GetAllTodos(c))
	s.Require().Equal(http.StatusOK, rec.Code)
}

func (s *TodosSuite) TestDeleteTodo() {
	reqStruct := TodoRequest{
		Title:          "TestTitle",
		Description:    "TestDescription",
		TimeToComplete: "2010-09-22 12:42:31",
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		s.T().Fail()
	}

	req := s.buildDeleteRequest(body)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todo/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c = s.addToken(c)

	h := NewTodosHandler(s.mockService)
	s.mockService.
		EXPECT().
		DeleteTodo(s.todoID).
		Return(nil)
	s.Require().NoError(h.DeleteTodo(c))
	s.Require().Equal(http.StatusNoContent, rec.Code)
}

func (s *TodosSuite) TestUpdateTodo() {
	reqStruct := TodoRequest{
		Title:          "TestTitle",
		Description:    "TestDescription",
		TimeToComplete: "2010-09-22 12:42:31",
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		s.T().Fail()
	}

	req := s.buildPatchRequest(body)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todo/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	c = s.addToken(c)

	h := NewTodosHandler(s.mockService)
	s.mockService.
		EXPECT().
		UpdateTodo(s.todoID, s.userID, reqStruct.Title, reqStruct.Description, s.parseTime(reqStruct.TimeToComplete)).
		Return(nil, nil)
	s.Require().NoError(h.UpdateTodo(c))
	s.Require().Equal(http.StatusNoContent, rec.Code)
}

func (s *TodosSuite) TestCreateTodo() {
	reqStruct := TodoRequest{
		Title:          "TestTitle",
		Description:    "TestDescription",
		TimeToComplete: "2010-09-22 12:42:31",
	}
	body, err := json.Marshal(reqStruct)
	if err != nil {
		s.T().Fail()
	}

	req := s.buildPostRequest(body)
	e := echo.New()
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/todo")

	c = s.addToken(c)

	h := NewTodosHandler(s.mockService)
	s.mockService.
		EXPECT().
		CreateTodo(s.userID, reqStruct.Title, reqStruct.Description, s.parseTime(reqStruct.TimeToComplete)).
		Return(nil, nil)
	s.Require().NoError(h.CreateTodo(c))
	s.Require().Equal(http.StatusCreated, rec.Code)
}

func TestTodosSuite(t *testing.T) {
	suite.Run(t, new(TodosSuite))
}
