package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("product not found")
	ErrCantDecodeProducts = errors.New("error to decode products")
	ErrUserIdNotValid     = errors.New("user id not valid")
	ErrCantUpdateUser     = errors.New("error to update user")
	ErrCantRemoveItemCart = errors.New("error to remove item cart")
	ErrCantGetItemCart    = errors.New("error to get item cart")
	ErrCantBuyCartItem    = errors.New("error to buy cart item")
)

func AddroductToCart() {

}

func RemoveCartIten() {

}
func BuyItenFromCart() {

}
func InstantBuy() {

}
