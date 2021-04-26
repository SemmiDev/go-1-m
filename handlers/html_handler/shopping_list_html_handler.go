package html_handler

import (
	"github.com/gin-gonic/gin"
	"grocery/models/entity"
	"grocery/repositories"
	"grocery/utils"
	"net/http"
	"strconv"
)

func ShowHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	shoppingList, _ := repositories.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "commons/not_found.tmpl", gin.H{})
	} else {
		c.HTML(http.StatusOK, "shopping_list/show.tmpl", shoppingList)
	}
}

func NewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "shopping_list/new.tmpl", gin.H{})
}

func CreateHandler(c *gin.Context) {
	var shoppingList entity.ShoppingList

	// App level validation
	bindErr := c.ShouldBind(&shoppingList)
	if bindErr != nil {
		shoppingList.Error = bindErr
		c.HTML(http.StatusOK, "shopping_list/new.tmpl", shoppingList)
		return
	}

	// Inserting data
	id, insertErr := repositories.Create(shoppingList)
	if insertErr != nil {
		c.HTML(http.StatusInternalServerError, "commons/internal_error.tmpl", gin.H{})
		utils.PanicError(insertErr)
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/show/"+strconv.FormatInt(id, 10))
	}
}

func EditHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	shoppingList, _ := repositories.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "commons/not_found.tmpl", gin.H{})
		return
	}

	c.HTML(http.StatusOK, "shopping_list/edit.tmpl", shoppingList)
}

func UpdateHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	var shoppingList entity.ShoppingList

	// App level validation
	bindErr := c.ShouldBind(&shoppingList)
	if bindErr != nil {
		shoppingList.Error = bindErr
		c.HTML(http.StatusOK, "shopping_list/edit.tmpl", shoppingList)
		return
	}

	foundShoppingList, _ := repositories.FindById(id)
	// Check if resource exist
	if foundShoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "commons/not_found.tmpl", gin.H{})
	}

	// Updating data
	shoppingList, updateErr := repositories.Put(foundShoppingList.Id, shoppingList)
	if updateErr != nil {
		c.HTML(http.StatusInternalServerError, "commons/internal_error.tmpl", gin.H{})
		utils.PanicError(updateErr)
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/show/"+strconv.FormatInt(id, 10))
	}
}

func IndexHandler(c *gin.Context) {
	limit, offset, page := utils.GetLimitOffset(c)
	shoppingLists := repositories.IndexWithPage(limit, offset)
	count := repositories.Count()
	pagination := utils.ProcessPagination("shopping-list", count, page, limit)

	m := make(map[string]interface{})
	m["shoppingLists"] = shoppingLists
	m["pagination"] = pagination

	c.HTML(http.StatusOK, "shopping_list/index.tmpl", m)
}

func DeleteHandler(c *gin.Context) {
	id := utils.GetInt64IdFromReqContext(c)
	shoppingList, _ := repositories.FindById(id)

	// Check if resource exist
	if shoppingList.Id == 0 {
		c.HTML(http.StatusNotFound, "commons/not_found.tmpl", gin.H{})
		return
	}

	err := repositories.Delete(shoppingList)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "commons/internal_error.tmpl", gin.H{})
		return
	} else {
		c.Redirect(http.StatusFound, "/shopping-list/")
	}
}