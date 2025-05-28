package main

import (
	"github.com/andresidrim/cesupa-hospital/database"
	pacientsHandler "github.com/andresidrim/cesupa-hospital/handlers/pacients"
	pacientsService "github.com/andresidrim/cesupa-hospital/services/pacients"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	pacientService := pacientsService.NewService(db)
	pacientHandler := pacientsHandler.NewHandler(pacientService)

	r := gin.Default()

	//  TODO: Configure CORS

	r.POST("/pacients", pacientHandler.AddPacient)
	r.GET("/pacients", pacientHandler.GetAllPacients)
	r.GET("/pacients/:id", pacientHandler.GetPacient)
	r.PUT("/pacients/:id", pacientHandler.UpdatePacient)
	r.DELETE("/pacients/:id", pacientHandler.DeletePacient)
	r.POST("/pacients/:id/appointment", pacientHandler.ScheduleAppointment)

	r.Run(":8080")
}
