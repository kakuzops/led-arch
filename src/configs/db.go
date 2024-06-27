package configs

import (
	"github.com/kakuzops/led-arch/src/models"
	"github.com/kakuzops/led-arch/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createdAccount() {
	db := utils.ConnectDB()

	users := [2]models.User{
		{Username: "kakuzops", Email: "kakuzops@kakuzops.com"},
		{Username: "Vascudagama", Email: "Vascudagama@kakuzops.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := utils.HashAndSalt([]byte(users[i].Username))
		user := models.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := models.BankAccount{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
}

func Migrate() {
	db := connectDb()
	db.AutoMigrate(&models.User{}, &models.BankAccount{})

	createdAccount()
}
