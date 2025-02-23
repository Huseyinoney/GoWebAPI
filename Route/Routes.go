package routes

import (
	handler "GoTutorial/Handler"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	userGroup := app.Group("/users") // Tüm user endpointleri için grup oluşturuyoruz.

	userGroup.Post("/add", userHandler.AddUser)                       // Kullanıcı ekleme
	userGroup.Delete("/delete", userHandler.DeleteUser)               // Kullanıcı silme
	userGroup.Put("/update-password", userHandler.UpdateUserPassword) // Şifre güncelleme
}
