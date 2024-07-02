package api

import (
	"ims/database"
	"ims/handlers"
	"ims/repositories"
	"ims/services"

	"github.com/gin-gonic/gin"
)

type server struct {
	listenAddr string
	db         database.Database
}

type Server interface {
	Start()
}

func NewServer(listenAddr string, db database.Database) *server {
	return &server{
		listenAddr: listenAddr,
		db:         db,
	}
}

func (s *server) Start() error {

	userRepository := repositories.NewUserRepository(s.db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()
	router.GET("/users", userHandler.HandleGetAllUsers)
	router.POST("/users", userHandler.HandleCreateUser)
	router.GET("/users/:id", userHandler.HandleGetUserByID)
	router.PUT("/users/:id", userHandler.HandleUpdateUser)
	router.DELETE("/users/:id", userHandler.HandleDeleteUser)
	return router.Run(s.listenAddr)
}
