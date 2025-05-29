package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/mocks"
	"github.com/andresidrim/cesupa-hospital/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// --- Helpers ---

func setupGetUserRouter(ms *mocks.MockUserService) *gin.Engine {
	h := NewHandler(ms)
	r := gin.Default()
	r.GET("/users/:id", h.GetUser)
	return r
}

func setupGetAllRouter(ms *mocks.MockUserService) *gin.Engine {
	h := NewHandler(ms)
	r := gin.Default()
	r.GET("/users", h.GetAllUsers)
	return r
}

// --- Test GetUser ---

func TestGetUserHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		paramID        string
		mockGet        func(id uint64) (*models.User, error)
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
			name:    "not found",
			paramID: "1",
			mockGet: func(id uint64) (*models.User, error) {
				return nil, assert.AnError
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "User not found",
		},
		{
			name:    "success",
			paramID: "42",
			mockGet: func(id uint64) (*models.User, error) {
				return &models.User{Model: gorm.Model{ID: 42}, Name: "Alice"}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `"name":"Alice"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &mocks.MockUserService{MockGet: tt.mockGet}
			router := setupGetUserRouter(ms)

			req := httptest.NewRequest("GET", "/users/"+tt.paramID, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}

// --- Test GetAllUsers ---

func TestGetAllUsersHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		query          string
		mockGetAll     func(roles []enums.Role) ([]models.User, error)
		expectedStatus int
		expectedLen    int
		expectedBody   string // optional substring to assert
	}{
		{
			name:           "no filter",
			query:          "",
			mockGetAll:     func(roles []enums.Role) ([]models.User, error) { return []models.User{{Name: "A"}}, nil },
			expectedStatus: http.StatusOK,
			expectedLen:    1,
			expectedBody:   `"name":"A"`,
		},
		{
			name:  "with roles",
			query: "?roles=doctor,admin",
			mockGetAll: func(roles []enums.Role) ([]models.User, error) {
				assert.ElementsMatch(t, roles, []enums.Role{enums.Doctor, enums.Admin})
				return []models.User{{Name: "B"}}, nil
			},
			expectedStatus: http.StatusOK,
			expectedLen:    1,
			expectedBody:   `"name":"B"`,
		},
		{
			name:           "service error",
			query:          "",
			mockGetAll:     func(roles []enums.Role) ([]models.User, error) { return nil, assert.AnError },
			expectedStatus: http.StatusNotFound,
			expectedLen:    0,
			expectedBody:   "No users were found",
		},
		{
			name:           "empty result",
			query:          "",
			mockGetAll:     func(roles []enums.Role) ([]models.User, error) { return []models.User{}, nil },
			expectedStatus: http.StatusOK,
			expectedLen:    0,
			expectedBody:   `"users":[]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := &mocks.MockUserService{MockGetAll: tt.mockGetAll}
			router := setupGetAllRouter(ms)

			req := httptest.NewRequest("GET", "/users"+tt.query, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			// Se esperamos OK, parseia JSON e confere len
			if w.Code == http.StatusOK {
				var body struct{ Users []models.User }
				err := json.Unmarshal(w.Body.Bytes(), &body)
				assert.NoError(t, err)
				assert.Len(t, body.Users, tt.expectedLen)
			}

			assert.Contains(t, w.Body.String(), tt.expectedBody)
		})
	}
}
