package main

import (
	"fmt"
	"learn-go-goroutine/config"
	"learn-go-goroutine/handler"
	"learn-go-goroutine/middleware"
	"learn-go-goroutine/models"
	"learn-go-goroutine/repo"
	"learn-go-goroutine/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// มีประโยชน์มั้ง
// https://reserved-poppy-sheep-762.medium.com/%E0%B8%A1%E0%B8%B2%E0%B8%97%E0%B8%B3-password-hashing-%E0%B8%94%E0%B9%89%E0%B8%A7%E0%B8%A2-bcrypt-%E0%B8%9E%E0%B8%A3%E0%B9%89%E0%B8%AD%E0%B8%A1%E0%B8%81%E0%B8%B1%E0%B8%9A%E0%B8%97%E0%B8%B3-jwt-%E0%B8%94%E0%B9%89%E0%B8%A7%E0%B8%A2-golang-%E0%B8%81%E0%B8%B1%E0%B8%99-db38fe2e4d38
// https://kritwis.medium.com/golang-%E0%B8%97%E0%B8%B3-jwt-%E0%B8%94%E0%B9%89%E0%B8%A7%E0%B8%A2-gin-framework-%E0%B9%81%E0%B8%A5%E0%B8%B0-go-module-%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B8%87%E0%B9%88%E0%B8%B2%E0%B8%A2-6e016fb8e30
// https://medium.com/@rluders/improving-request-validation-and-response-handling-in-go-microservices-cc54208123f2

var validate *validator.Validate

func main() {
	app := fiber.New()

	db := initDb()

	validate = validator.New(validator.WithRequiredStructEnabled())

	userRepo := repo.NewUserRepoDb(db)
	userService := service.NewUserService(userRepo,validate)
	userHandler := handler.NewUserHttphandler(userService)
	
	authRepo := repo.NewAuthRepoDb(db)
	authServive := service.NewAuthService(authRepo,validate)
	authHandler := handler.NewAuthHttpHandler(authServive)

	app.Post("/user", userHandler.Register)
	app.Get("/user", middleware.Protected(),userHandler.GetUsers)
	
	app.Post("/signIn",authHandler.SignIn)


	app.Listen(":3000")
}  

func initDb() *gorm.DB{

	dsn := fmt.Sprintf(
	"host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
	config.Config("DB_PORT"),
	config.Config("DB_USER"),
	config.Config("DB_PASSWORD"),
	config.Config("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}
	
	db.AutoMigrate(&models.User{})

	return db
}