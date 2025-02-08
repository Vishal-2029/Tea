package controllers

import (
	"fmt"
	"strconv"
	db "teamaking/config"
	"teamaking/models"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	TeaName string `json:"tea_name"`
	Type    string `json:"type"`
}

func CreateTea(c *fiber.Ctx) error {
	var body RequestBody

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}

	Tea := models.Tea{
		TeaName: body.TeaName,
		Type:    body.Type,
	}
	db.DB.Create(&Tea)
	fmt.Println("tea", body)
	return c.Status(200).JSON(fiber.Map{
		"Success": true,
		"Message": "Tea is ready",
		"Data":    Tea,
	})
}

type Teas struct {
	Id      uint   `json:"Id"`
	TeaName string `json:"TeaName"`
}

func ReadTea(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var Tea []Teas
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&Tea).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	TeaData := map[string]interface{}{
		"Tea":  Tea,
		"meta": metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    TeaData,
	})
}

func UpdateTea(c *fiber.Ctx) error {
	TeaId := c.Params("TeaID") // Corrected parameter name
	var Tea models.Tea

	if err := db.DB.First(&Tea, TeaId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Tea Not Found",
		})
	}

	var updateTeaData models.Tea
	if err := c.BodyParser(&updateTeaData); err != nil || updateTeaData.TeaName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Valid tea name is required",
		})
	}

	Tea.TeaName = updateTeaData.TeaName;
	Tea.Type = updateTeaData.Type
	db.DB.Save(&Tea)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Update Tea successfully",
		"data":    Tea,
	})
}

func DeleteTea(c *fiber.Ctx) error {
	TeaId := c.Params("TeaID") // Corrected parameter name
	var Tea models.Tea

	if err := db.DB.First(&Tea, TeaId).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Tea Not Found",
		})
	}

	db.DB.Delete(&Tea)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Tea Deleted successfully",
	})
}
