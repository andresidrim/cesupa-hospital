package pacients

import (
	"net/http"
	"strconv"

	"github.com/andresidrim/cesupa-hospital/models"
	ps "github.com/andresidrim/cesupa-hospital/services/pacients"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type Handler struct {
	service ps.PacientService
}

func NewHandler(service ps.PacientService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) AddPacient(c *gin.Context) {
	var payload AddPacientDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input: " + err.Error()})
		return
	}

	var pacient models.Pacient
	if err := copier.Copy(&pacient, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to copy data: " + err.Error()})
		return
	}

	if err := h.service.Create(&pacient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create pacient" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"pacient": pacient})
}

func (h *Handler) GetPacient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	pacient, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Pacient not found: " + err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"pacient": pacient})
}

func (h *Handler) GetAllPacients(c *gin.Context) {
	name := c.Query("name")
	ageStr := c.Query("age")

	pacients, err := h.service.GetAll(name, ageStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No pacient was found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pacients": pacients})
}
