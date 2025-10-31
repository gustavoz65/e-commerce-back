package database

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavoz65/e-commerce-back/database"
)

var (
	ErrCantFindProduct    = errors.New("product not found")
	ErrCantDecodeProducts = errors.New("error decoding products")
	ErrUserIdNotValid     = errors.New("user id is not valid")
	ErrCantUpdateUser     = errors.New("error updating user")
	ErrCantRemoveItemCart = errors.New("error removing item from cart")
	ErrCantGetItemCart    = errors.New("error getting item from cart")
	ErrCantBuyCartItem    = errors.New("error buying cart item")
)

func AddProductToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("userID")
		product_id := c.Query("productID")

		if user_id == "" || product_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "userID and product ID are required"})
			c.Abort()
			return
		}
		err := database.AddProductToCart(user_id, product_id)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "falied internal")
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"message": "product added to cart successfully"})
	}
}

func RemoveCartItem() {

}
func BuyItemFromCart() {

}
func InstantBuy() {

}
