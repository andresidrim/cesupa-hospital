package users

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/andresidrim/cesupa-hospital/enums"
	us "github.com/andresidrim/cesupa-hospital/services/users"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service us.UserService
}

func NewHandler(service us.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	user, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	rawRoles := c.Query("roles")
	var roles []enums.Role
	if rawRoles != "" {
		for _, role := range strings.Split(rawRoles, ",") {
			roles = append(roles, enums.Role(strings.TrimSpace(role)))
		}
	}

	users, err := h.service.GetAll(roles)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users were found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *Handler) GetDoctors(c *gin.Context) {
	// força o filtro de papel "doctor"
	users, err := h.service.GetAll([]enums.Role{enums.Doctor})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No doctors found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctors": users})
}
