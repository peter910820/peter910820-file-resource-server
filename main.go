package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is index page!")
	})

	logrus.Fatal(app.Listen(":3000"))
}
