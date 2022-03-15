package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ismael3s/alura/go/05/database"
	"github.com/ismael3s/alura/go/05/models"
)

func ShowStudents(c *gin.Context) {
	var students []models.Student

	result := database.DB.Find(&students)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, students)
}

func ShowStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var student models.Student

	result := database.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"messagse": result.Error.Error(),
		})
		return
	}

	c.JSON(200, student)

}

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	result := database.DB.Create(&student)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(201, student)

}

func RemoveStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var student models.Student

	result := database.DB.Delete(&student, id)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func UpdateStundet(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	var student models.Student
	var studentDB models.Student

	if err = c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if student.ID != uint(id) {
		c.JSON(400, gin.H{
			"message": "id diferente do id informado",
		})
		return
	}

	database.DB.First(&studentDB, id)

	database.DB.Model(&studentDB).Updates(student)

	c.JSON(200, student)
}
