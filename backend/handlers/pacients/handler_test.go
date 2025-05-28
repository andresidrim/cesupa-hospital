package pacients

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andresidrim/cesupa-hospital/mocks"
	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestRouter(h *Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/pacients", h.GetAllPacients)
	return router
}

func TestAddPacient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		body           string
		mockCreateErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "valid input",
			body: `{
				"name": "John Doe",
				"birthDate": "2000-01-01T00:00:00Z",
				"cpf": "12345678900",
				"sex": "male",
				"phoneNumber": "+123456789",
				"address": "123 Street"
			}`,
			mockCreateErr:  nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   "John Doe",
		},
		{
			name: "missing required field",
			body: `{
				"birthDate": "2000-01-01T00:00:00Z"
			}`,
			mockCreateErr:  nil, // Won't be called
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid input",
		},
		{
			name:           "invalid JSON",
			body:           `{name: "John"}`, // invalid JSON
			mockCreateErr:  nil,              // Won't be called
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid input",
		},
		{
			name: "service error",
			body: `{
				"name": "John Doe",
				"birthDate": "2000-01-01T00:00:00Z",
				"cpf": "12345678900",
				"sex": "male",
				"phoneNumber": "+123456789",
				"address": "123 Street"
			}`,
			mockCreateErr:  assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to create pacient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.MockPacientService{
				MockCreate: func(pacient *models.Pacient) error {
					t.Logf("MockCreate called with: %+v", pacient)
					return tt.mockCreateErr
				},
			}

			handler := NewHandler(mockService)
			router := gin.Default()
			router.POST("/pacients", handler.AddPacient)

			req, _ := http.NewRequest(http.MethodPost, "/pacients", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			t.Logf("Response: %s", resp.Body.String())
			t.Log("END OF SUBTEST")
			t.Log(strings.Repeat("-", 50))

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.Contains(t, resp.Body.String(), tt.expectedBody)
		})
	}
}

func TestGetAllPacients(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		query      string
		mockResult []models.Pacient
		mockError  error
		wantStatus int
		wantBody   string
	}{
		{
			name:       "found pacient",
			query:      "?name=John&age=30",
			mockResult: []models.Pacient{{Name: "John Doe"}},
			wantStatus: http.StatusOK,
			wantBody:   "John Doe",
		},
		{
			name:       "no pacients",
			query:      "?name=Unknown",
			mockResult: []models.Pacient{},
			wantStatus: http.StatusOK,
			wantBody:   "[]",
		},
		{
			name:       "internal error",
			query:      "?name=Error",
			mockError:  assert.AnError,
			wantStatus: http.StatusNotFound,
			wantBody:   "No pacient was found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.MockPacientService{
				MockGetAll: func(name, ageStr string) ([]models.Pacient, error) {
					t.Logf("MockGetAll called with name: %s, age: %s", name, ageStr)
					return tt.mockResult, tt.mockError
				},
			}

			handler := NewHandler(mockService)
			router := setupTestRouter(handler)

			req, _ := http.NewRequest(http.MethodGet, "/pacients"+tt.query, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)
			t.Log("Response Body:", resp.Body.String())
			t.Log(strings.Repeat("-", 50))

			assert.Equal(t, tt.wantStatus, resp.Code)
			assert.Contains(t, resp.Body.String(), tt.wantBody)
		})
	}
}

func TestUpdatePacient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		paramID        string
		body           string
		mockGetErr     error
		mockUpdateErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid ID",
			paramID:        "abc",
			body:           `{}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid ID",
		},
		{
			name:           "pacient not found",
			paramID:        "1",
			body:           `{ "name": "John", "birthDate": "2000-01-01T00:00:00Z", "cpf":"123", "sex":"male", "phoneNumber":"123", "address":"street" }`,
			mockGetErr:     assert.AnError,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Pacient not found",
		},
		{
			name:           "invalid input",
			paramID:        "1",
			body:           `{ "name": 123 }`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid Input",
		},
		{
			name:           "update error",
			paramID:        "1",
			body:           `{ "name": "John", "birthDate": "2000-01-01T00:00:00Z", "cpf":"123", "sex":"male", "phoneNumber":"123", "address":"street" }`,
			mockGetErr:     nil,
			mockUpdateErr:  assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to update pacient",
		},
		{
			name:           "successful update",
			paramID:        "1",
			body:           `{ "name": "John", "birthDate": "2000-01-01T00:00:00Z", "cpf":"123", "sex":"male", "phoneNumber":"123", "address":"street" }`,
			mockGetErr:     nil,
			mockUpdateErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "pacient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.MockPacientService{
				MockGet: func(id uint64) (*models.Pacient, error) {
					return nil, tt.mockGetErr
				},
				MockUpdate: func(id uint64, pacient *models.Pacient) error {
					return tt.mockUpdateErr
				},
			}

			handler := NewHandler(mockService)
			router := gin.Default()
			router.PUT("/pacients/:id", handler.UpdatePacient)

			req, _ := http.NewRequest(http.MethodPut, "/pacients/"+tt.paramID, bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			t.Logf("Response: %s", resp.Body.String())
			t.Log("END OF SUBTEST")
			t.Log(strings.Repeat("-", 50))

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.Contains(t, resp.Body.String(), tt.expectedBody)
		})
	}
}

func TestDeletePacient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		paramID        string
		mockGetErr     error
		mockDeleteErr  error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid ID",
			paramID:        "abc",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid ID",
		},
		{
			name:           "pacient not found",
			paramID:        "1",
			mockGetErr:     gorm.ErrRecordNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Pacient not found",
		},
		{
			name:           "delete failure",
			paramID:        "1",
			mockGetErr:     nil,
			mockDeleteErr:  assert.AnError,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to delete pacient",
		},
		{
			name:           "successful delete",
			paramID:        "1",
			mockGetErr:     nil,
			mockDeleteErr:  nil,
			expectedStatus: http.StatusOK,
			expectedBody:   "pacient",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mocks.MockPacientService{
				MockGet: func(id uint64) (*models.Pacient, error) {
					if tt.mockGetErr != nil {
						return nil, tt.mockGetErr
					}
					return &models.Pacient{
						Model: gorm.Model{
							ID: uint(id),
						},
						Name: "John Doe",
					}, nil
				},
				MockDelete: func(id uint64) error {
					return tt.mockDeleteErr
				},
			}

			handler := NewHandler(mockService)
			router := gin.Default()
			router.DELETE("/pacients/:id", handler.DeletePacient)

			req, _ := http.NewRequest(http.MethodDelete, "/pacients/"+tt.paramID, nil)
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			t.Logf("Response: %s", resp.Body.String())
			t.Log("END OF SUBTEST")
			t.Log(strings.Repeat("-", 50))

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.Contains(t, resp.Body.String(), tt.expectedBody)
		})
	}
}
