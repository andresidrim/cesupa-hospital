package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/middlewares"
	"github.com/andresidrim/cesupa-hospital/mocks"
	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/andresidrim/cesupa-hospital/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// setup router for Register
func setupRegisterRouter(ms *mocks.MockAuthService) *gin.Engine {
	h := NewHandler(ms)
	r := gin.Default()
	r.POST("/register", h.Register)
	return r
}

// setup router for Login
func setupLoginRouter(ms *mocks.MockAuthService) *gin.Engine {
	h := NewHandler(ms)
	r := gin.Default()
	r.POST("/login", h.Login)
	return r
}

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name            string
		body            string
		mockRegisterErr error
		expectedStatus  int
		expectedBody    string
	}{
		{
			name:           "invalid input",
			body:           `{ "name": 123 }`, // fails binding
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid input",
		},
		{
			name:            "service error",
			body:            `{ "name":"Alice", "cpf":"111", "password":"secret", "role":"admin" }`,
			mockRegisterErr: assert.AnError,
			expectedStatus:  http.StatusInternalServerError,
			expectedBody:    "Failed to register user",
		},
		{
			name:            "success",
			body:            `{ "name":"Alice", "cpf":"111", "password":"secret", "role":"doctor" }`,
			mockRegisterErr: nil,
			expectedStatus:  http.StatusCreated,
			expectedBody:    `"name":"Alice"`, // JSON includes created user
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &mocks.MockAuthService{
				MockRegister: func(u *models.User) error {
					u.ID = 99
					return tt.mockRegisterErr
				},
			}
			r := setupRegisterRouter(ms)

			req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		body           string
		mockToken      string
		mockLoginErr   error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid input",
			body:           `{ "cpf": "123" }`, // missing password
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid input",
		},
		{
			name:           "auth failure",
			body:           `{ "cpf":"123", "password":"wrong" }`,
			mockLoginErr:   assert.AnError,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   assert.AnError.Error(),
		},
		{
			name:           "success",
			body:           `{ "cpf":"123", "password":"secret" }`,
			mockToken:      "tok123",
			mockLoginErr:   nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `"token":"tok123"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &mocks.MockAuthService{
				MockLogin: func(cpf, pass string) (string, error) {
					return tt.mockToken, tt.mockLoginErr
				},
			}
			r := setupLoginRouter(ms)

			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

func TestRegisterHandlerUnauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUserSvc := &mocks.MockUserService{
		MockGet: func(id uint64) (*models.User, error) {
			return &models.User{Model: gorm.Model{ID: uint(id)}, Role: enums.Doctor}, nil
		},
	}

	mockAuthSvc := &mocks.MockAuthService{}

	r := gin.Default()
	protected := r.Group("/")
	protected.Use(
		middlewares.JWTAuthMiddleware(mockUserSvc),
		middlewares.RoleMiddleware(enums.Admin),
	)
	protected.POST("/register", NewHandler(mockAuthSvc).Register)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/register", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	token, _ := utils.GenerateJWT(1)
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/register", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusForbidden, w.Code)
}
