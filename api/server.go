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

	roleRepository := repositories.NewRoleRepository(s.db)
	roleService := services.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	router.GET("/roles", roleHandler.HandleGetAllRoles)
	router.POST("/roles", roleHandler.HandleCreateRole)
	router.GET("/roles/:id", roleHandler.HandleGetRoleByID)
	router.PUT("/roles/:id", roleHandler.HandleUpdateRole)
	router.DELETE("/roles/:id", roleHandler.HandleDeleteRole)

	userRoleRepository := repositories.NewUserRoleRepository(s.db)
	userRoleService := services.NewUserRoleService(userRoleRepository, userRepository)
	userRoleHandler := handlers.NewUserRoleHandler(userRoleService)

	router.POST("user_roles", userRoleHandler.HandleCreateUserRole)
	router.DELETE("user_roles", userRoleHandler.HandleDeleteUserRole)
	router.GET("user_roles/:role_id", userRoleHandler.HandleGetRoleMembers)

	accessRepository := repositories.NewAccessRepository(s.db)
	accessService := services.NewAccessService(accessRepository)
	accessHandler := handlers.NewAccessHandler(accessService)

	router.GET("/accesses", accessHandler.HandleGetAllAccesses)
	router.POST("/accesses", accessHandler.HandleCreateAccess)
	router.GET("/accesses/:id", accessHandler.HandleGetAccessByID)
	router.DELETE("/accesses/:id", accessHandler.HandleDeleteAccess)
	router.GET("/accesses/name", accessHandler.HandleGetAccessByNameRequest)
	router.GET("/accesses/code", accessHandler.HandleGetAccessByCodeRequest)

	roleAccessRepository := repositories.NewRoleAccessRepository(s.db)
	roleAccessService := services.NewRoleAccessService(roleAccessRepository, accessRepository)
	roleAccessHandler := handlers.NewRoleAccessHandler(roleAccessService)

	router.POST("role_accesses", roleAccessHandler.HandleCreateRoleAccess)
	router.DELETE("role_accesses", roleAccessHandler.HandleDeleteRoleAccess)
	router.GET("role_accesses/:role_id", roleAccessHandler.HandleGetAccessesByRole)

	return router.Run(s.listenAddr)
}
