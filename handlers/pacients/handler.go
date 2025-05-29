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

// AddPacient cria um novo paciente
// @Summary      Cadastra um novo paciente
// @Description  Registra um paciente com dados obrigatórios e opcionais
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Param        paciente  body      AddPacientDTO  true  "Dados do paciente"
// @Success      201       {object}  models.Pacient
// @Failure      400       {object}  gin.H         "Invalid input"
// @Failure      500       {object}  gin.H         "Failed to create pacient"
// @Router       /pacients [post]
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

// GetPacient busca um paciente pelo ID
// @Summary      Busca paciente
// @Description  Retorna os dados de um paciente pelo seu ID
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do paciente"
// @Success      200  {object}  models.Pacient
// @Failure      400  {object}  gin.H        "Invalid ID"
// @Failure      404  {object}  gin.H        "Pacient not found"
// @Router       /pacients/{id} [get]
func (h *Handler) GetPacient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	pacient, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Pacient not found: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pacient": pacient})
}

// GetAllPacients lista pacientes com filtros opcionais
// @Summary      Lista pacientes
// @Description  Retorna todos os pacientes, podendo filtrar por nome e/ou idade
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Param        name  query     string  false  "Filtra pelo nome (substring)"
// @Param        age   query     int     false  "Filtra pela idade exata"
// @Success      200   {array}   models.Pacient
// @Failure      404   {object}  gin.H        "No pacient was found"
// @Router       /pacients [get]
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

// UpdatePacient altera dados de um paciente
// @Summary      Atualiza paciente
// @Description  Atualiza campos de um paciente existente (PATCH semantics)
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Param        id        path     int               true  "ID do paciente"
// @Param        paciente  body     UpdatePacientDTO  true  "Dados que serão atualizados"
// @Success      200       {object} models.Pacient
// @Failure      400       {object} gin.H        "Invalid ID or Input"
// @Failure      404       {object} gin.H        "Pacient not found"
// @Failure      500       {object} gin.H        "Failed to update pacient"
// @Router       /pacients/{id} [put]
func (h *Handler) UpdatePacient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	if _, err := h.service.Get(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Pacient not found: " + err.Error()})
		return
	}

	var payload UpdatePacientDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Input: " + err.Error()})
		return
	}

	var updatedPacient models.Pacient
	if err := copier.Copy(&updatedPacient, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to copy data: " + err.Error()})
		return
	}

	if err := h.service.Update(id, &updatedPacient); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update pacient: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pacient": updatedPacient})
}

// DeletePacient remove logicamente um paciente
// @Summary      Exclui paciente
// @Description  Inativa (ou exclui) um paciente pelo ID
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID do paciente"
// @Success      200  {object}  models.Pacient
// @Failure      400  {object}  gin.H        "Invalid ID"
// @Failure      404  {object}  gin.H        "Pacient not found"
// @Failure      500  {object}  gin.H        "Failed to delete pacient"
// @Router       /pacients/{id} [delete]
func (h *Handler) DeletePacient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	deletedPacient, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Pacient not found: " + err.Error()})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete pacient: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pacient": deletedPacient})
}

// ScheduleAppointment agenda uma consulta para um paciente
// @Summary      Agenda consulta
// @Description  Cria uma nova consulta para o paciente informado
// @Tags         Pacientes
// @Accept       json
// @Produce      json
// @Param        id          path      int                  true  "ID do paciente"
// @Param        appointment  body     ScheduleAppointmentDTO true  "Dados da consulta"
// @Success      201         {object}  models.Appointment
// @Failure      400         {object}  gin.H              "Invalid ID or Input"
// @Failure      404         {object}  gin.H              "Pacient not found"
// @Failure      500         {object}  gin.H              "Failed to create appointment"
// @Router       /pacients/{id}/appointments [post]
func (h *Handler) ScheduleAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID: " + err.Error()})
		return
	}

	if _, err := h.service.Get(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Pacient not found: " + err.Error()})
		return
	}

	var payload ScheduleAppointmentDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input: " + err.Error()})
		return
	}

	var appointment models.Appointment
	if err := copier.Copy(&appointment, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to copy data: " + err.Error()})
		return
	}

	appointment.PacientID = uint(id)

	if err := h.service.ScheduleAppointment(&appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create appointment: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"appointment": appointment})
}
