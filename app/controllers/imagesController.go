package controllers

import (
	"github.com/GlebSolncev/http-rest-api/app/models"
	"github.com/GlebSolncev/http-rest-api/app/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ImagesController struct {
	// ...
}

func (controller *ImagesController) Index(c *gin.Context) {
	var images []models.Image
	_ = services.GetAllImages(&images)

	c.JSON(http.StatusOK, images)
}

func (controller *ImagesController) Store(c *gin.Context) {
	var image models.Image
	c.BindJSON(&image)
	err := services.CreateImages(&image)

	if err != nil {
		log.Fatal(err)
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, image)
	}
}
