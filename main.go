package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/e-commerce-back/controllers"
	"github.com/gustavoz65/e-commerce-back/database"
	"github.com/gustavoz65/e-commerce-back/middleware"
	"github.com/gustavoz65/e-commerce-back/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Product"), database.UserData(database.Client), "users")
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartCheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(" " + port))

}
