package controller

import (
	"github.com/edmaciel/go-fiber-fx/internal/domain"
	"github.com/gofiber/fiber/v2"
)

func FindUser(ctx *fiber.Ctx) error {
	user := domain.User{Name: "User"}
	ctx.JSON(user)

	return nil
}
