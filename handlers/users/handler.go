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

// GetUser retorna um usuário pelo ID
// @Summary      Busca usuário
// @Description  Retorna dados de um usuário pelo seu ID
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do usuário"
// @Success      200  {object}  models.User
// @Failure      400  {object}  ErrorResponse       "Invalid ID"
// @Failure      404  {object}  ErrorResponse       "User not found"
// @Router       /users/{id} [get]
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

// GetAllUsers lista usuários, opcionalmente filtrados por papel(s)
// @Summary      Lista usuários
// @Description  Retorna todos os usuários, podendo filtrar por um ou mais papéis
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Param        roles  query     []string  false  "Filtro de papéis separados por vírgula"
// @Success      200    {array}   models.User
// @Failure      404    {object}  ErrorResponse      "No users were found"
// @Router       /users [get]
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

// GetDoctors retorna apenas os usuários com papel de médico
// @Summary      Lista médicos
// @Description  Retorna todos os usuários cujo papel é 'doctor'
// @Tags         Usuários
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.User
// @Failure      404  {object}  ErrorResponse      "No doctors found"
// @Router       /doctors [get]
func (h *Handler) GetDoctors(c *gin.Context) {
	// força o filtro de papel "doctor"
	users, err := h.service.GetAll([]enums.Role{enums.Doctor})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No doctors found: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctors": users})
}
