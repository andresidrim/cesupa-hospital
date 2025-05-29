package auth

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresidrim/cesupa-hospital/mocks"
	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(ms *mocks.MockAuthService) *gin.Engine {
	h := NewHandler(ms)
	router := gin.Default()
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	return router
}

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name     string
		body     string
		mockErr  error
		status   int
		contains string
	}{
		{
			name:     "invalid JSON",
			body:     `{invalid`,
			status:   http.StatusBadRequest,
			contains: "Invalid input",
		},
		{
			name:     "missing field",
			body:     `{"cpf":"123","password":"secret","role":"admin"}`, // falta name
			status:   http.StatusBadRequest,
			contains: "Invalid input",
		},
		{
			name:     "service error",
			body:     `{"name":"A","cpf":"123","password":"secret","role":"admin"}`,
			mockErr:  assert.AnError,
			status:   http.StatusInternalServerError,
			contains: "Failed to register user",
		},
		{
			name:     "success",
			body:     `{"name":"A","cpf":"123","password":"secret","role":"doctor"}`,
			status:   http.StatusCreated,
			contains: `"name":"A"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &mocks.MockAuthService{
				MockRegister: func(user *models.User) error {
					return tt.mockErr
				},
			}
			router := setupRouter(ms)

			req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
			assert.Contains(t, w.Body.String(), tt.contains)
		})
	}
}

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name      string
		body      string
		mockToken string
		mockErr   error
		status    int
		contains  string
	}{
		{
			name:     "invalid JSON",
			body:     `{invalid`,
			status:   http.StatusBadRequest,
			contains: "Entrada inválida",
		},
		{
			name:     "user not found",
			body:     `{"cpf":"x","password":"y"}`,
			mockErr:  assert.AnError,
			status:   http.StatusUnauthorized,
			contains: "not found",
		},
		{
			name:     "incorrect password",
			body:     `{"cpf":"x","password":"y"}`,
			mockErr:  assert.AnError, // same error path
			status:   http.StatusUnauthorized,
			contains: "incorrect",
		},
		{
			name:      "success",
			body:      `{"cpf":"x","password":"y"}`,
			mockToken: "tok123",
			status:    http.StatusOK,
			contains:  `"token":"tok123"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &mocks.MockAuthService{
				MockLogin: func(cpf, pass string) (string, error) {
					return tt.mockToken, tt.mockErr
				},
			}
			router := setupRouter(ms)

			req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
			assert.Contains(t, w.Body.String(), tt.contains)
		})
	}
}
