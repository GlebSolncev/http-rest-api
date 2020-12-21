package routers

import (
	"github.com/GlebSolncev/http-rest-api/app/controllers"
	"github.com/gin-gonic/gin"
)

//Route{"SaveCustomers", "POST", "/customer", controller.ImagesController.Index}

func RouteList() *gin.Engine {
	res := gin.Default()
	imageController := controllers.ImagesController{}
	image := res.Group("/images")
	{
		image.GET(
			"/index",
			imageController.Index,
		)

		image.POST(
			"/",
			imageController.Store,
		)
	}

	return res
}
