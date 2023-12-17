package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var (
	counter int = 0
	images []string = []string{"image.jpg", "image1.jpg", "image2.png"}
	comp int = 0
	pages []string
)

func main() {
	updatePages()
	app := fiber.New()
	app.Static("/", "./public")
	app.Get("/sum", func(c *fiber.Ctx) error {
		counter++
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("<h1 class='red' style='font-family: Arial, sans-serif;' id='swapper'>"+strconv.Itoa(counter)+"</h1>")
	})
	app.Get("/sub", func(c *fiber.Ctx) error {
		counter--
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("<h1 class='blue' style='font-family: Arial, sans-serif;' id='swapper'>"+strconv.Itoa(counter)+"</h1>")
	})
	app.Get("/init", func(c *fiber.Ctx) error {
		return c.SendString(strconv.Itoa(counter))
	})
	app.Get("/search", func(c *fiber.Ctx) error {
		tekst := c.Query("query")
		return c.SendString("Searching: " + tekst)
	})
	app.Get("/images", func(c *fiber.Ctx) error {
		which :=  c.QueryInt("which", 0) % len(images)
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.SendString("<img width=1000 src='/" + images[which] + "' />")
	})
	app.Listen(":3000")
}

func updatePages() {
	files, err := filepath.Glob("./public/arrows/*.html")
	if err != nil {
		panic(err)
	}
	pages = files
	fmt.Println(pages)
}
