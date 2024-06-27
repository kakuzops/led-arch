package utils

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}

}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

func ConnectDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgresql://postgres:postgres@localhost:5432/ldarc"), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	HandleErr(err)
	return db
}
