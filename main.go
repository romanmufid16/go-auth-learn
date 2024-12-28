package main

import (
	"github.com/gin-gonic/gin"
	"github.com/romanmufid16/go-auth-learn/config"
	"github.com/romanmufid16/go-auth-learn/controllers"
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
	}

	r.Run()
}
