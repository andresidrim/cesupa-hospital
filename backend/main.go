package main

import (
	"github.com/andresidrim/cesupa-hospital/database"
	"github.com/andresidrim/cesupa-hospital/env"
	"github.com/gin-contrib/cors"

	authHandlers "github.com/andresidrim/cesupa-hospital/handlers/auth"
	pacientsHandler "github.com/andresidrim/cesupa-hospital/handlers/pacients"
	usersHandler "github.com/andresidrim/cesupa-hospital/handlers/users"

	authServices "github.com/andresidrim/cesupa-hospital/services/auth"
	pacientsService "github.com/andresidrim/cesupa-hospital/services/pacients"
	usersService "github.com/andresidrim/cesupa-hospital/services/users"

	"github.com/andresidrim/cesupa-hospital/enums"
	"github.com/andresidrim/cesupa-hospital/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexão ao banco
	db := database.Connect()

	// Services
	pacientSvc := pacientsService.NewService(db)
	userSvc := usersService.NewService(db)
	authSvc := authServices.NewService(db)

	// Handlers
	pacientH := pacientsHandler.NewHandler(pacientSvc)
	userH := usersHandler.NewHandler(userSvc)
	authH := authHandlers.NewHandler(authSvc)

	// Middlewares
	jwtMw := middlewares.JWTAuthMiddleware(userSvc)
	roleAdmin := middlewares.RoleMiddleware(enums.Admin)
	roleRecepAdmin := middlewares.RoleMiddleware(enums.Receptionist, enums.Admin)
	roleRecepDoctor := middlewares.RoleMiddleware(enums.Receptionist, enums.Doctor)

	// Setup Gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Rota pública de login
	r.POST("/login", authH.Login)

	// Tudo que vier a seguir exige JWT
	authGroup := r.Group("/")
	authGroup.Use(jwtMw)
	{
		// Registro só por Admin
		authGroup.POST("/register",
			roleAdmin,
			authH.Register,
		)

		// Listar médicos (para agendamento) → Recepcionist ou Admin
		authGroup.GET("/doctors",
			roleRecepAdmin,
			userH.GetDoctors,
		)

		// 1. Cadastrar novo paciente → Recepcionist ou Admin
		authGroup.POST("/pacients",
			roleRecepAdmin,
			pacientH.AddPacient,
		)

		// 2. Consultar dados de paciente → Recepcionist ou Doctor
		authGroup.GET("/pacients",
			roleRecepDoctor,
			pacientH.GetAllPacients,
		)
		authGroup.GET("/pacients/:id",
			roleRecepDoctor,
			pacientH.GetPacient,
		)

		// 3. Atualizar dados de paciente → Recepcionist ou Admin
		authGroup.PUT("/pacients/:id",
			roleRecepAdmin,
			pacientH.UpdatePacient,
		)

		// 4. Inativar cadastro de paciente → Recepcionist ou Admin
		authGroup.DELETE("/pacients/:id",
			roleRecepAdmin,
			pacientH.DeletePacient,
		)

		// 5. Agendamento de consulta → Recepcionist ou Admin
		authGroup.POST("/pacients/:id/appointment",
			roleRecepAdmin,
			pacientH.ScheduleAppointment,
		)

		// Gestão de usuários (listar e consultar) → apenas Admin
		authGroup.GET("/users",
			roleAdmin,
			userH.GetAllUsers,
		)
		authGroup.GET("/users/:id",
			roleAdmin,
			userH.GetUser,
		)
	}

	// Start server
	r.Run(":" + env.PORT)
}
