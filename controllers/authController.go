package controllers

import (
	"fmt"
	"go_auth/database"
	"go_auth/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "sdasda_123sadqw"

func HiApi(c *fiber.Ctx) error {
	return c.SendString("Hi, api in Golang with Gorm and Fiber V2")
}

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 8)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	database.DB.Create(&user)
	return c.JSON(&user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		fmt.Println("error: ", err)
		return err
	}
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	clams := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, e := clams.SignedString([]byte(secretKey))
	if e != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
		"token":   token,
		"user":    user,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, e := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if e != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Invalida token",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)
	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
