package controller

import (
	"demo/database"
	"demo/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Claims struct {
	jwt.StandardClaims
}

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	if data["password"] != data["confirm_password"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password doesn't match",
		})
	}
	encodedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  encodedPassword,
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "User Not found",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		return c.JSON(fiber.Map{
			"message": "email or password incorrect",
		})
	}
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"jwt": token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	//fmt.Print(token)
	if err != nil || !token.Valid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*Claims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	fmt.Print(claims.Issuer)
	return c.JSON(user)
}
