package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"

	"github.com/sirupsen/logrus"
)

func main() {
	os.MkdirAll("./image", os.ModePerm)
	// logrus settings
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	// Create a new engine
	engine := jet.New("./views", ".jet")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/public", "./public")
	app.Static("/image", "./image")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "資源伺服器目錄",
		}, "layouts/base")
	})

	app.Get("/folder", func(c *fiber.Ctx) error {
		dir := "./image"
		fileName := []string{}
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				logrus.Error(err)
				return err
			}
			if !info.IsDir() {
				fileName = append(fileName, path)
			}
			return nil
		})
		if err != nil {
			logrus.Error(err)
			return c.Status(fiber.StatusInternalServerError).SendString("server has an error")
		}
		return c.Render("folder", fiber.Map{
			"FileName": fileName,
		}, "layouts/base")
	})

	app.Post("/api/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			logrus.Error(err)
			return c.Status(fiber.StatusBadRequest).SendString("uploaded failed")
		}
		err = c.SaveFile(file, fmt.Sprintf("./image/%s", file.Filename))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("save file failed")
		}
		return c.Status(fiber.StatusOK).SendString("save file successful!")
	})

	logrus.Fatal(app.Listen(":3000"))
}
