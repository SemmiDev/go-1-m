package api_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grocery/models/entity"
	"grocery/repositories"
	"grocery/utils"
	"net/http"
)

func CreateDosenPAHandler(c *gin.Context) {
	var dosenPA entity.DosenPA

	// App level validation
	bindErr := c.BindJSON(&dosenPA)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := repositories.CreateDosenPA(dosenPA)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		utils.PanicError(insertErr)
	} else {
		dosenPA.Id = id
		c.JSON(http.StatusCreated, dosenPA)
	}
}

func ShowDosenPAHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	dosenPA, _ := repositories.FindDosenPAById(id)

	// Check if resource exist
	if dosenPA.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, dosenPA)
	}
}

func PutDosenPAHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	var dosenPA entity.DosenPA

	// App level validation
	bindErr := c.BindJSON(&dosenPA)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	// Check if resource exist
	foundDosenPA, _ := repositories.FindDosenPAById(id)
	if foundDosenPA.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Updating data
	std, err := repositories.PutDosenPAByID(foundDosenPA.Id, dosenPA)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusCreated, std)
	}
}

func DeleteDosenPAHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)

	// Check if resource exist
	foundDosenPA, _ := repositories.FindDosenPAById(id)
	if foundDosenPA.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Deleting data
	err := repositories.DeleteDosenByID(foundDosenPA)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusNoContent, "Successful Deletion")
	}
}