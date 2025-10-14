package controllers

import "github.com/gin-gonic/gin"

func HasPassword(oassword string) string {

}

func CheckPassword(userpassword string, givenpassword string) (bool, string) {

}

func Singup() gin.HandlerFunc {
  return func(c * gin.Context) {
	 var ctx, cancel = context.WhithTimeout(context.Background(), 100*time.Second)
	 defer cancel()

	 var user models.User

	 c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
		return
	 }

	 validationErr := Validade.Struct(user)
	 if validationErr := nil {
		c.JSON(http.StatusBadRequest, ginH("error": validationErr))
		return
	 }

	count, err := UserCollection.CountDocuments(ctx, bson.M())
     if err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
		return
	 }

	 if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H("error": "user already exists"))
	 }

	 count, err = UserCollection.CountDocuments(ctx, bson.M("email": user.Email))
     if err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
		return
	 }
  }
}

func Login() gin.HandlerFunc {

}

func ProductViewAdmin() gin.HandlerFunc {

}

func Searchproduct() gin.HandlerFunc {

}

func SearchproductByQuery() gin.HandlerFunc {

}
