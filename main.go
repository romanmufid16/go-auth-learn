package main

import (
	"github.com/gin-gonic/gin"
	"github.com/romanmufid16/go-auth-learn/config"
	"github.com/romanmufid16/go-auth-learn/controllers"
	"github.com/romanmufid16/go-auth-learn/middleware"
	"github.com/romanmufid16/go-auth-learn/repository"
	"github.com/romanmufid16/go-auth-learn/service"
	"github.com/romanmufid16/go-auth-learn/utils"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func init() {
	utils.LoadEnvHandler()
	config.SetupDatabaseConnection()
	config.SyncDatabase()
}

func main() {
	r := gin.Default()

	db := config.GetDB()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", userController.RegisterController)
		authRoutes.POST("/login", userController.LoginController)
	}

	protected := r.Group("api/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/current", userController.GetUserInfo)
		protected.PUT("/update", userController.UpdateUser)
		protected.DELETE("/:id/delete", userController.DeleteUser)
	}

	r.Run()
}
