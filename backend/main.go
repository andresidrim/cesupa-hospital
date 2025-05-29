package main

import (
	"github.com/andresidrim/cesupa-hospital/database"
	"github.com/andresidrim/cesupa-hospital/env"
	pacientsHandler "github.com/andresidrim/cesupa-hospital/handlers/pacients"
	usersHandler "github.com/andresidrim/cesupa-hospital/handlers/users"
	pacientsService "github.com/andresidrim/cesupa-hospital/services/pacients"
	usersService "github.com/andresidrim/cesupa-hospital/services/users"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()

	pacientService := pacientsService.NewService(db)
	pacientHandler := pacientsHandler.NewHandler(pacientService)

	userService := usersService.NewService(db)
	userHandler := usersHandler.NewHandler(userService)

	r := gin.Default()

	//  TODO: Configure CORS

	r.POST("/pacients", pacientHandler.AddPacient)
	r.GET("/pacients", pacientHandler.GetAllPacients)
	r.GET("/pacients/:id", pacientHandler.GetPacient)
	r.PUT("/pacients/:id", pacientHandler.UpdatePacient)
	r.DELETE("/pacients/:id", pacientHandler.DeletePacient)
	r.POST("/pacients/:id/appointment", pacientHandler.ScheduleAppointment)
	r.GET("users/:id", userHandler.GetUser)
	r.GET("users", userHandler.GetAllUsers)

	r.Run(":" + env.PORT)
}
