package handler

import (
	"learn-go-goroutine/service"
	"learn-go-goroutine/types"

	"github.com/gofiber/fiber/v2"
)

type userHttpHandler struct {
	service service.UserService
}

func NewUserHttphandler(service service.UserService) *userHttpHandler{
	return &userHttpHandler{service:service}
}

func (u *userHttpHandler) Register(c *fiber.Ctx) error{
	newUser := new(types.User)

	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	result,err := u.service.Register(newUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (u *userHttpHandler) GetUsers(c *fiber.Ctx) error{
	users,err := u.service.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(users) 
}