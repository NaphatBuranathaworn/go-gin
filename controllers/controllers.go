package controllers

import (
	"net/http"

	"github.com/NaphatBuranathaworn/go-gin/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    var todo []models.Todo
    c.BindJSON(&todo)

    err := models.GetAllTodos(&todo)

    if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func CreateATodo(c *gin.Context) {
    var todo models.Todo
    c.BindJSON(&todo)

    err := models.CreateATodo(&todo)
    if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {

}

func UpdateATodo(c *gin.Context) {

}

func DeleteATodo(c *gin.Context) {

}