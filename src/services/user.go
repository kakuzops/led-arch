package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kakuzops/led-arch/src/interfaces"
	"github.com/kakuzops/led-arch/src/models"
	"github.com/kakuzops/led-arch/src/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {

	db := utils.ConnectDb()
	user := &interfaces.User{}

	if db.Where("username = ?", username).First(&user).Error != nil {
		return map[string]interface{}{"message": "User not found"}
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}

	accounts := []models.ResponseAccount{}
	db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

	responseUser := &models.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	defer db.Close()

	tokenContet := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContet)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	utils.HandleErr(err)

	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response
}
