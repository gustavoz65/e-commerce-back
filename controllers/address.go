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
	return func(c *gin.Context) {
		
		userIDStr := c.Query("user_id")
		if userIDStr == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusBadRequest, gin.H{"error": "user id is missing"})
			c.Abort()
			return
		}

		var address models.Address
		if err := c.BindJSON(&address); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		
		filter := bson.D{{Key: "_id", Value: userID}}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "address.$[i].house_name", Value: address.House},
				{Key: "address.$[i].street_name", Value: address.Street},
				{Key: "address.$[i].city_name", Value: address.City},
				{Key: "address.$[i].pin_code", Value: address.Pincode},
			}},
		}

		arrayFilters := options.ArrayFilters{
			Filters: []interface{}{
				bson.M{"i.address_type": "work"},
			},
		}
		updateOptions := options.UpdateOptions{
			ArrayFilters: &arrayFilters,
		}

		_, err = UserCollection.UpdateOne(ctx, filter, update, &updateOptions)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to update address", "error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"success": "address updated successfully"})
	}
}



func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H("ërror": "user id is missing"))
			c.Abort()
			return
		}

		address := make([]models.Address, 0)
		user_id, err := primitive.ObjectIDFromHex(user_id)
		if err != nil {
			c.IndentedJSON(http.StatusBadResquest, "Internal server error")
		}
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		filter := bson.D{primitive.E{key: "_id", Value: user_id}}   // chaves de id de atualização do mongodb
		update := bson.D{{key: "$set", Value: bson.D{primitive.E{key: "address", Value: address}}}} // seta o campo de endereço como uma matriz vazia

		_, err := UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, "message": "wordng command")
		}
		defer cancel()
		ctx.Done()
		c.IndentedJSON(http.StatusOK, "success": "address deleted successfully")
	}
}
