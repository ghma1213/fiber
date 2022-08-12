package controllers

import (
	"fiber/src/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	type Admin struct {
		Admin_Seq int
		Admin_Id  string
		Admin_Nm  string
	}
	var admins []Admin

	db := database.DBConn
	db.Raw("SELECT admin_seq, admin_id, admin_nm FROM admin_member").Scan(&admins)

	data := fiber.Map{"admin": admins}
	fmt.Println(data)
	return c.Render("index", data)
}
