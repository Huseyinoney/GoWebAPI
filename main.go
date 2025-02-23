package main

import (
	"GoTutorial/Database"
	handler "GoTutorial/Handler"
	model "GoTutorial/Model"
	repository "GoTutorial/Repository"
	routes "GoTutorial/Route"
	service "GoTutorial/Service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := Database.ConnectDB()

	// Veritabanı migrasyonlarını yap
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("Migrasyon başarısız: ", err)
	}
	// Servis katmanını oluştur
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// Kullanıcı rotalarını ayarla
	routes.SetupUserRoutes(app, userHandler)

	// Sunucuyu başlat
	app.Listen(":3000")
}
