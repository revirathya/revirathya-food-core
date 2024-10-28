package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revirathya/revirathya-food-core/internals/modules"
)

type RouteConfig struct {
	App     *fiber.App
	Usecase *modules.Usecase
}
