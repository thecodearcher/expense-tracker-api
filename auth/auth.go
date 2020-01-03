package auth

import (
	"expense-tracker/database"
	"expense-tracker/helpers"
	"expense-tracker/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"golang.org/x/crypto/bcrypt"
)

// Login authenticates user into application
func Login(w http.ResponseWriter, req *http.Request) {
	res := &helpers.APIResponse{}
	var loginData = new(models.LoginUser)
	schema.NewDecoder().Decode(loginData, req.PostForm)
	err := loginData.Validate()
	if err != nil {
		res.Data = err
		res.Message = err.Error()
		res.Error(http.StatusBadRequest, w)
		return
	}
	var user models.User
	result := database.DB.Table("users").Where(&models.User{Email: loginData.Email}).Find(&user)
	if result.RecordNotFound() {
		err = result.Error
		goto authErr
	}

	if result.Error != nil {
		res.Data = err
		res.Error(http.StatusInternalServerError, w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))

authErr:
	if err != nil {
		res = &helpers.APIResponse{
			Data:    loginData,
			Message: "Invalid email or password entered!",
		}
		res.Error(http.StatusUnauthorized, w)
		return
	}

	token := generateJWT(&user)
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	res = &helpers.APIResponse{
		Data: data,
	}
	res.Success(w)
	return
}

// Signup authenticates user into application
func Signup(w http.ResponseWriter, req *http.Request) {
	res := &helpers.APIResponse{}

	user := new(models.User)
	decoder := schema.NewDecoder()
	decoder.Decode(user, req.PostForm)
	err := user.Validate()
	if err != nil {
		res.Data = err
		res.Message = err.Error()
		res.Error(http.StatusBadRequest, w)
		return
	}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(passwordHash)
	result := database.DB.Create(&user)
	if result.Error != nil {
		res.Data = err
		res.Error(http.StatusInternalServerError, w)
		return
	}
	token := generateJWT(user)
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	res = &helpers.APIResponse{
		Data: data,
	}
	res.Created(w)
	return
}

// Generates signed JWT and returns token or panics on error
func generateJWT(user *models.User) string {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}

	return tokenString
}
