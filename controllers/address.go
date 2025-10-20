package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/e-commerce-back/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {

}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
      return func (c *gin.Context) {
		  user_id := c.Query("id")

		  if user_id == "" {
            c.Header("Contect-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H("ërror": "user id is missing"))
			c.Abort()
			return 
		  }
		  address := make([]models.Address, 0)
		  usert_id, err := primitive.ObjectIDFromHex(user_id)
	      if err != nil {
			c.IndentedJSON(hhtp.StatusBadResquest, "Internal server error")
		  }
		  ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		  defer cancel()

		  filter := bson.D{primitive.E{key: "_id", Value: usert_id}}   // chaves de id de atualização do mongodb
		  update := bson.D{{key: "$set", Value: bson.D{primitive.E{key: "address", Value: address}}}} // seta o campo de endereço como uma matriz vazia
		
		  _, err := UserCollection.UpdateOne(ctx,filter, update)
		  if err != nil {
			c.IndentedJSON(http.StatusBadRequest,"message": "wordng command")
		  }
		  defer cancel()
		  ctx.Done()
		  c.IndentedJSON(http.StatusOK, "success": "address deleted successfully")
		}
}
