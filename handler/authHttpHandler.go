package handler

import (
	"learn-go-goroutine/service"
	"learn-go-goroutine/types"

	"github.com/gofiber/fiber/v2"
)

type authHttpHandler struct {
	service service.AuthService
}

func NewAuthHttpHandler(service service.AuthService) *authHttpHandler{
	return &authHttpHandler{service:service}
}

func (a *authHttpHandler) SignIn(c *fiber.Ctx) error{
	cred := new(types.SignInRequest)

	if err := c.BodyParser(cred); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	result,err := a.service.SignIn(cred)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}