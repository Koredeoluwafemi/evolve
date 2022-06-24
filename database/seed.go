package database

import (

	//"golang.org/x/crypto/bcrypt"
	"evolve/models"
	"gorm.io/gorm"
)

func seed(db *gorm.DB) {

	//user
	//userSeeder(db)

}

func userSeeder(db *gorm.DB) {

	//user := models.User{
	//	Firstname: "ABC",
	//	Lastname:  "EFG",
	//	Email:     "123@ev.com",
	//	Phone:     "0801",
	//	Address:   "adeyeye street",
	//}
	//db.FirstOrCreate(&user)

	//user := models.User{
	//	Firstname: "Wing",
	//	Lastname:  "Hurley",
	//	Email:     "a@protonmail.ca",
	//	Phone:     "0802",
	//	Address:   "adeyeye street",
	//}
	//db.Create(&user)

	user := models.User{
		Firstname: "Mia",
		Lastname:  "Miranda",
		Email:     "bibendum.fermentum@aol.com",
		Phone:     "0803",
		Address:   "adeyeye street",
	}
	db.Create(&user)
}
