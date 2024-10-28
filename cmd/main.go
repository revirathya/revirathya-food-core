package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/revirathya/revirathya-food-core/internals/config"
	"github.com/revirathya/revirathya-food-core/internals/modules"
)

func main() {
	// Read Config
	C, err := config.LoadConfig("../envs")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	// Init App
	app := fiber.New(fiber.Config{
		AppName: "revirathya-finance-core",
	})
	err = modules.Bootstrap(C, app)
	if err != nil {
		log.Fatal("Bootstrap failed: ", err)
	}
}
