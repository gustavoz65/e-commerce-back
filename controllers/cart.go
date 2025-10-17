package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}

}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
   poductQueryID := c.Query("id")
   if poductQueryID == "" { 
	log.Println("product id is empty")
	_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty")))
	return
}

  if userQueryID == "" {
	log.Println("user id is emoty")

	_ = c.AbortWithError(http.StatusBadRequest, error.New("user id is empty"))
	return
  }

  productID, err:= primitive.ObjectIDFromHex(poductQueryID)
   if err != nil {
	log.Println("product id is not valid")
    c.AbortWithError(http.StatusBadRequest)
    return 
}

   var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

   defer cancel()

  err =  database.AddroductToCart(ctx,app.userCollection, app.prodCollection, productID, userQueryID)
  if err != nil {
	c.InsertAbortWithError(http.StatusInternalServerError, err)
	return
  }	
  c.IndentedJSON(http.StatusCreated, "message": "success")
}
}

func RemoveItem() gin.HandlerFunc {

}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func InstantBuy() gin.HandlerFunc {

}
