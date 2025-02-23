package handler

import (
	model "GoTutorial/Model"
	service "GoTutorial/Service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Kullanıcı Ekleme
func (userHandler *UserHandler) AddUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if success := userHandler.userService.AddUser(user.Name, user.Surname, user.Address, user.Age, user.Password); !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User added successfully",
	})
}

// Kullanıcı Silme
func (userHandler *UserHandler) DeleteUser(c *fiber.Ctx) error {
	var requestData struct {
		ID uint64 `json:"id"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if success := userHandler.userService.DeleteUser(requestData.ID); !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

// Kullanıcı Şifre Güncelleme
func (userHandler *UserHandler) UpdateUserPassword(c *fiber.Ctx) error {
	var requestData struct {
		ID       uint64 `json:"id"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	if success := userHandler.userService.UpdateUserPassword(requestData.ID, requestData.Password); !success {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password updated successfully",
	})
}
