package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"project/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
}

type MockContext struct {
	mock.Mock
	echo.Context
}

func (m *MockContext) Request() *http.Request {
	args := m.Called()
	return args.Get(0).(*http.Request)
}

func (m *MockContext) Response() *httptest.ResponseRecorder {
	args := m.Called()
	return args.Get(0).(*httptest.ResponseRecorder)
}

func (m *MockContext) Param(name string) string {
	args := m.Called(name)
	return args.String(0)
}

func (m *MockContext) Bind(i interface{}) error {
	args := m.Called(i)
	return args.Error(0)
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	mockArgs := m.Called(query, args)
	return mockArgs.Get(0).(*gorm.DB)
}

// make unit testing
func TestGetActivityLogsController(t *testing.T) {

	eError := echo.New()
	mockDBError := new(MockDB)
	expectedActivityLogsError := []model.ActivityLog{}
	fmt.Println(expectedActivityLogsError)
	mockDBError.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{Error: fmt.Errorf("error")})

	// Create a new HTTP request
	reqError := httptest.NewRequest(http.MethodGet, "/activity-logs", nil)
	recError := httptest.NewRecorder()
	cError := eError.NewContext(reqError, recError)

	// Call the handler function
	handlerError := GetActivityLogsController(mockDBError)
	errError := handlerError(cError)

	assert.Nil(t, errError)
	assert.Equal(t, http.StatusInternalServerError, recError.Code)

	expectedResponseError := echo.Map{
		"code":    "500",
		"message": "error",
	}
	jsonStringError, errErrorJsonString := toJSONString(expectedResponseError)
	assert.Nil(t, errErrorJsonString)
	assert.JSONEq(t, recError.Body.String(), jsonStringError)

	mockDBError.AssertCalled(t, "Find", mock.Anything, mock.Anything)

	e := echo.New()
	mockDB := new(MockDB)
	expectedActivityLogs := []model.ActivityLog{}
	fmt.Println(expectedActivityLogs)
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{RowsAffected: 0})

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/activity-logs", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler function
	handler := GetActivityLogsController(mockDB)
	err := handler(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	expectedResponse := echo.Map{
		"code":    "404",
		"message": "Data Not Found",
	}
	jsonString, err := toJSONString(expectedResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	mockDB.AssertCalled(t, "Find", mock.Anything, mock.Anything)

	e = echo.New()
	mockDB = new(MockDB)
	expectedActivityLogs = []model.ActivityLog{}
	fmt.Println(expectedActivityLogs)
	mockDB.On("Find", mock.Anything, mock.Anything).Return(&gorm.DB{RowsAffected: 1})

	req = httptest.NewRequest(http.MethodGet, "/activity-logs", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	handler = GetActivityLogsController(mockDB)
	err = handler(c)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expectedResponse = echo.Map{
		"code":    "200",
		"message": "success get quiz",
		"data":    expectedActivityLogs,
	}

	jsonString, err = toJSONString(expectedResponse)
	assert.Nil(t, err)
	assert.JSONEq(t, rec.Body.String(), jsonString)

	mockDB.AssertCalled(t, "Find", mock.Anything, mock.Anything)
}

func toJSONString(v interface{}) (string, error) {
	// Mengonversi objek menjadi JSON string
	jsonString, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}
