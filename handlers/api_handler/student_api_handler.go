package api_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grocery/models/entity"
	"grocery/repositories"
	"grocery/utils"
	"net/http"
)

func CreateStudentHandler(c *gin.Context) {
	var student entity.Student

	// App level validation
	bindErr := c.BindJSON(&student)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := repositories.CreateStudent(student)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		utils.PanicError(insertErr)
	} else {
		student.Id = id
		c.JSON(http.StatusCreated, student)
	}
}

func ShowStudentHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	student, _ := repositories.FindStudentById(id)

	// Check if resource exist
	if student.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func PutStudentHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	var student entity.Student

	// App level validation
	bindErr := c.BindJSON(&student)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	// Check if resource exist
	foundStudent, _ := repositories.FindStudentById(id)
	if foundStudent.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Updating data
	std, err := repositories.PutStudentByID(foundStudent.Id, student)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusCreated, std)
	}
}

func DeleteStudentHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)

	// Check if resource exist
	foundStudent, _ := repositories.FindStudentById(id)
	if foundStudent.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Deleting data
	err := repositories.DeleteStudentByID(foundStudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusNoContent, "Successful Deletion")
	}
}

func ShowStudentJoinDOsenHandler(c *gin.Context) {
	student, _ := repositories.StudentJoinDosenPA()
	c.JSON(http.StatusOK, student)
}