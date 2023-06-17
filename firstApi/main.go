package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Students struct {
	Id     int     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	SName  string  `json:"fullname"`
	Course string  `json:"course"`
	Fee    float64 `json:"fee"`
}

func getStudent(c *gin.Context) {

	var student []Students
	DB.Find(&student)

	c.IndentedJSON(http.StatusOK, student)

}

func addStudent(c *gin.Context) {

	var newStudent Students

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	student := Students{Id: newStudent.Id, SName: newStudent.SName, Course: newStudent.Course, Fee: newStudent.Fee}
	DB.Create(&student)
	c.IndentedJSON(http.StatusCreated, student)

}

func findStudent(c *gin.Context)  {

	var newStudent Students

	if err := DB.Where("id = ?", c.Param("id")).First(&newStudent).Error; err !=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Record not Found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, newStudent)

	
}

func updateStudent(c *gin.Context)  {
	var newStudent Students

	if err := DB.Where("id = ?", c.Param("id")).First(&newStudent).Error; err !=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Record not Found!"})
		return
	}

	var supdate Students

	if err := c.ShouldBindJSON(&supdate); err !=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	DB.Model(&newStudent).Updates(supdate)
	c.IndentedJSON(http.StatusOK, newStudent)

	
}

func deleteStudent(c *gin.Context)  {

	var newStudent Students

	if err := DB.Where("id = ?", c.Param("id")).First(&newStudent).Error; err !=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Record not Found!"})
		return
	}

	DB.Delete(&newStudent)

	c.IndentedJSON(http.StatusOK, newStudent)

	
}


var Data *gorm.DB

func main() {

	router := gin.Default()
	conDb()
	router.GET("/students", getStudent)
	router.POST("/students", addStudent)
	router.GET("/students/:id", findStudent)
	router.PATCH("/students/:id", updateStudent)
	router.DELETE("/student/:id", deleteStudent)
	router.Run("localhost:1234")

}
