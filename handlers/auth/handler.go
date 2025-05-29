package auth

import (
	"net/http"

	"github.com/andresidrim/cesupa-hospital/models"
	as "github.com/andresidrim/cesupa-hospital/services/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service as.AuthService
}

func NewHandler(service as.AuthService) *Handler {
	return &Handler{service: service}
}

// Register godoc
// @Summary     Cadastra um novo usuário
// @Description Recebe name, cpf, password e role e cria o usuário
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       payload body     RegisterDTO true "Dados para registro"
// @Success     201     {object} handlers.RegisterResponse
// @Failure     400     {object} handlers.ErrorResponse
// @Failure     500     {object} handlers.ErrorResponse
// @Router      /register [post]
func (h *Handler) Register(c *gin.Context) {
	var payload RegisterDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input: " + err.Error()})
		return
	}

	user := models.User{
		Name:     payload.Name,
		CPF:      payload.CPF,
		Password: payload.Password,
		Role:     payload.Role,
	}

	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":   user.ID,
			"name": user.Name,
			"cpf":  user.CPF,
			"role": user.Role,
		},
	})
}

// Login godoc
// @Summary     Faz login e retorna JWT
// @Description Recebe cpf e senha e devolve um token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       payload body     LoginDTO true "Dados para login"
// @Success     200     {object} handlers.TokenResponse
// @Failure     400     {object} handlers.ErrorResponse
// @Failure     401     {object} handlers.ErrorResponse
// @Router      /login [post]
func (h *Handler) Login(c *gin.Context) {
	var payload LoginDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	token, err := h.service.Login(payload.CPF, payload.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
