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

	user1 := models.User{
		Firstname: "ABC",
		Lastname:  "EFG",
		Email:     "123@ev.com",
		Phone:     "0801",
		Address:   "adeyeye street",
	}
	db.Create(&user1)

	user2 := models.User{
		Firstname: "Wing",
		Lastname:  "Hurley",
		Email:     "a@protonmail.ca",
		Phone:     "0802",
		Address:   "adeyeye street",
	}
	db.Create(&user2)
	//
	user3 := models.User{
		Firstname: "Mia",
		Lastname:  "Miranda",
		Email:     "bibendum.fermentum@aol.com",
		Phone:     "0803",
		Address:   "adeyeye street",
	}
	db.Create(&user3)

	user4 := models.User{
		Firstname: "Buffy",
		Lastname:  "Evans",
		Email:     "evans@aol.com",
		Phone:     "08034",
		Address:   "adeyeye4 street",
	}
	db.Create(&user4)

	user5 := models.User{
		Firstname: "Skyler",
		Lastname:  "Bender",
		Email:     "bender@aol.com",
		Phone:     "08035",
		Address:   "adeyeye5 street",
	}
	db.Create(&user5)

	user6 := models.User{
		Firstname: "Burton",
		Lastname:  "Todd",
		Email:     "todd@aol.com",
		Phone:     "08036",
		Address:   "adeyeye6 street",
	}
	db.Create(&user6)

	user7 := models.User{
		Firstname: "Kameko",
		Lastname:  "Tate",
		Email:     "tate@aol.com",
		Phone:     "08037",
		Address:   "adeyeye7 street",
	}
	db.Create(&user7)

	user8 := models.User{
		Firstname: "Willa",
		Lastname:  "Dickson",
		Email:     "dikson@aol.com",
		Phone:     "08038",
		Address:   "adeyey8e street",
	}
	db.Create(&user8)

	user9 := models.User{
		Firstname: "Kyra",
		Lastname:  "Gutierrez",
		Email:     "kyra@aol.com",
		Phone:     "08039",
		Address:   "adeye9ye street",
	}
	db.Create(&user9)

	user10 := models.User{
		Firstname: "Sasha",
		Lastname:  "Frederick",
		Email:     "rederick@aol.com",
		Phone:     "080310",
		Address:   "adeyey10e street",
	}
	db.Create(&user10)

	user11 := models.User{
		Firstname: "Kuame",
		Lastname:  "Hensley",
		Email:     "kuame@aol.com",
		Phone:     "08011",
		Address:   "adeyeye11 street",
	}
	db.Create(&user11)
}
