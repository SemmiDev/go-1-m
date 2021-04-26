package api_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grocery/models/entity"
	"grocery/repositories"
	"grocery/utils"
	"net/http"
)

func CreateHandler(c *gin.Context) {
	var shoppingList entity.ShoppingList

	// App level validation
	bindErr := c.BindJSON(&shoppingList)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
		return
	}

	// Inserting data
	id, insertErr := repositories.Create(shoppingList)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Something wrong on our server"))
		utils.PanicError(insertErr)
	} else {
		shoppingList.Id = id
		c.JSON(http.StatusCreated, shoppingList)
	}
}

func ShowHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	shoppingList, _ := repositories.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
	} else {
		c.JSON(http.StatusOK, shoppingList)
	}
}

func PutHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	var shoppingList entity.ShoppingList

	// App level validation
	bindErr := c.BindJSON(&shoppingList)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprint(bindErr))
	}

	// Check if resource exist
	foundShoppingList, _ := repositories.FindById(id)
	if foundShoppingList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Updating data
	shoppingList, err := repositories.Put(foundShoppingList.Id, shoppingList)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusCreated, shoppingList)
	}
}

func DeleteHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)

	// Check if resource exist
	foundShoppingList, _ := repositories.FindById(id)
	if foundShoppingList.Id == 0 {
		c.JSON(http.StatusNotFound, "Not found")
		return
	}

	// Deleting data
	err := repositories.Delete(foundShoppingList)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusNoContent, "Successful Deletion")
	}
}
