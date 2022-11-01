package controllers

import (
	"fiber/src/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	type admin_members struct {
		gorm.Model
		Admin_Seq int    `gorm:"column:admin_seq"`
		Admin_Id  string `gorm:"column:admin_id"`
		Admin_Nm  string `gorm:"column:admin_nm"`
	}
	var admin_member []admin_members

	db := database.DBConn
	// db.Raw("SELECT admin_seq, admin_id, admin_nm FROM admin_member").Scan(&admins)
	db.Find(&admin_member)

	data := fiber.Map{"admin": admin_member}
	fmt.Println(data)
	return c.Render("index", data)
}
