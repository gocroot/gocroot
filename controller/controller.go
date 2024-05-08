package controller

import (
	"gocroot/helper"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("Hello Go Croot!!!")
}

func GetIPServer(c *fiber.Ctx) error {
	ipaddr := helper.GetIPaddress()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ipaddress": ipaddr})
}
