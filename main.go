package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
type User struct {
	Name string
}

func main() {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})

	if err != nil {
		panic("Cloud not connect to database ⚠️")
	}
	fmt.Print(db)
	app := fiber.New()

	app.Get("/", home)
	log.Fatal(app.Listen(":3000"))
}
func home(c *fiber.Ctx) error {
	return c.SendString("Hello xxx")
}
