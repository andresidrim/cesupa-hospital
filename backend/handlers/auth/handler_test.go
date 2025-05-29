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

func TestRegisterHandler_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// MockAuthService não usado no middleware, apenas para satisfazer assinatura
	ms := &mocks.MockAuthService{}

	r := gin.Default()
	// aplica autenticação e autorização como no main
	r.Use(
		middlewares.JWTAuthMiddleware(ms),
		middlewares.RoleMiddleware(enums.Admin),
	)
	r.POST("/register", NewHandler(ms).Register)

	// 1) Sem token
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/register", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// 2) Token válido mas sem role Admin -> 403
	// cria token para usuário de papel Doctor
	tok, _ := utils.GenerateJWT(42)
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/register", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusForbidden, w.Code)
}
