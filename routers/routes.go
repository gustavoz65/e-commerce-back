package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/e-commerce-back/controllers"
)

func UserRoutes(incomingRoutes *gin.Enginer) {
	incomingRoutes.POST("/users/singup", controllers.Singup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewAdmin())
	incomingRoutes.GET("/users/productsview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
}
