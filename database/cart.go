package database

import "errors"

var (
    ErrCantFindProduct    = errors.New("product not found")
    ErrCantDecodeProducts = errors.New("error decoding products")
    ErrUserIdNotValid     = errors.New("user id is not valid")
    ErrCantUpdateUser     = errors.New("error updating user")
    ErrCantRemoveItemCart = errors.New("error removing item from cart")
    ErrCantGetItemCart    = errors.New("error getting item from cart")
    ErrCantBuyCartItem    = errors.New("error buying cart item")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}
func BuyItemFromCart() {

}
func InstantBuy() {

}
