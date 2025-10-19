package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gustavoz65/e-commerce-back/controllers"
	"github.com/gustavoz65/e-commerce-back/database"
	"github.com/gustavoz65/e-commerce-back/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var UserCollecction *mongo.Collection = database.UserData(database.Client, "User")
var ProdCollection *mongo.Collection = database.UserData(database.Client, "Products")
var Validate = validator.New()


func HasPassword(password string) string {
            bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
            if err != nil {
                log.Panic(err)
            }
            return string(bytes)
}

func CheckPassword(userpassword string, givenpassword string) (bool, string) {
    err := bcrypt.CompareHashAndPassword([]byte(userpassword), []byte(givenpassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "password is not valid"
		valid = false
	}
	return valid, msg
}


func Signup() gin.HandlerFunc {
  return func(c * gin.Context) {
	 var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	 defer cancel()

	 var user models.User

	 c.BindJSON(&user); err != nil {
	 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	 }

	 validationErr := Validate.Struct(user)
	 if validationErr := nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	 }

	count, err := UserCollection.CountDocuments(ctx, bson.M("email":user.Email))
     if err != nil {
		log.Panic(err)
		c.JSON(http.StatusBadRequest, gin.H("error": err.Error()))
		return
	 }

	 if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
	 }

	 count, err := UserCollection.CountDocuments(ctx, bson.M("phone":user.Phone))

	 defer cancel()
     if err := nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H("error": err))
	 }

	 if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H("error": "phone number already exists"))
	 }
      
	 password := HashPassword(*user.Password)

    user.Password = &password

	 user.Created_At,_ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_ID = user.ID.Hex()
    token, refreshtoken, _ := GenerateAllTokens(*user.Email,*user.First_Name,*user.Last_Name,user.User_ID)
	user.Token = &token
	user.Refresh_Token = &refreshtoken
	 user.UserCart = make([]models.ProductUser{}, 0)
	user.Address_Details = make([]models.Address, 0)
	user.Order_Status = make([]models.Order, 0)
	_, insertErr := UserCollection.InsertOne(ctx, user)
	insertErr = nil {
		c.JSON(http.StatusInternalServerError, gin.H("error": "the user did not get created"))
        return
	}
	defer conect()

   
	c.JSON(http.StatusCreated, gin.H{"message": "success"})
  }
}

func Login() gin.HandlerFunc {
  return func (c *gin.Context) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()
  
 var user models.User

 if er := c.BindJSON(&user) , err != nil {
   c.JSON(http.StatusBadRequest, gin.H("error": err))
   return
 }

 err := UserCollection.FindOne(ctx, bson("email": user.Email)).Decode(*founduser)
 defer cancel()

 if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "email or password is incorrect"})
	return
 }

 passwordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
 defer cancel()
 
 if !passwordIsValid {
	c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	fmt.Println(msg)
	return
 }

 token, refrashtoken, _ := GenerateAllTokens(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)
 defer cancel()

 generate.UpdateAllTokens(token, refrashtoken, founduser.User_ID)

 .JSON(http.StatusFound, founduser)
}
}

func ProductViewAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandlerFunc {
	return func(c *gijn.Context) {
		var listaProduct []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
        defer cancel()

		productCursor, err := ProdCollection.Find(ctx, bson.D{{}}) 
		if err != nil {
			c.IdentedJSON(http.StatusInternalServerError, "something went wrong. please try after some time")
			return
		}
		err = productCursor.All(ctx, &listaProduct)
		if err != nile {
			log.Println(err)
			c.AbortWhithStatus(http.StatusInternalServerError)
			return         
		}

		defer productCursor.Close()

		if err := productCursor.Err(); err != nil {
			log.Println(err)
			c.IdentedJSON(400, "invalid request")
			return
		}
		defer cancel()

		c.IdentedJSON(http.StatusOK, listaProduct)
	}

}

func SearchProductByQuery() gin.HandlerFunc {

}
