package main

import (
	"evolve/config"
	"evolve/database"
	"evolve/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"math"
	"os"
	"strconv"
)

func main() {

	app := fiber.New()

	//initialize database
	database.Start()
	database.Migrate()

	app.Get("users", Users)

	port := config.App.Port

	if config.App.ENV == config.Environment.Stage {
		port = os.Getenv("PORT")
	}

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}

func Users(c *fiber.Ctx) error {

	db := database.DB

	type linkHolder struct {
		Link string
	}

	type userContainer struct {
		ID        uint   `json:"id"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
	}

	var (
		user    models.User
		users   []models.User
		links   []linkHolder
		link    linkHolder
		err     error
		dbCount int
	)

	if c.Query("email") != "" {
		dbCount = user.GetEmailCount(db, c.Query("email"))
	}

	//if c.Query("dr") != "" {
	//dbCount = user.GetDateRangeCount(db)
	//}

	if c.Query("email") == "" && c.Query("dr") == "" {
		dbCount = user.GetCount(db)
	}

	pageSize := 2

	if c.Query("pagesize") != "" {
		pageSize, err = strconv.Atoi(c.Query("pagesize"))
		if err != nil {
			output := fiber.Map{
				"status":  false,
				"message": "invalid page size supplied",
				"data":    "",
			}
			return c.Status(400).JSON(output)
		}
	}

	//get number of pages
	pagesNumber := int(math.Ceil(float64(dbCount) / float64(pageSize)))

	for i := 0; i < pagesNumber; i++ {
		link.Link = "users?page=" + strconv.Itoa(i+1)
		links = append(links, link)
	}

	getUsers := db.Scopes(Paginate(c, pageSize))
	if c.Query("email") != "" {
		getUsers.Where("email = ?", c.Query("email"))
	}
	getUsers.Order("id asc")
	rows := getUsers.Find(&users)
	if rows.RowsAffected == 0 {
		empty := make([]string, 0)

		output := fiber.Map{
			"status":  false,
			"message": "no data found",
			"data":    empty,
		}
		return c.Status(200).JSON(output)
	}

	allUsers := make([]userContainer, rows.RowsAffected)

	for i, item := range users {
		allUsers[i] = userContainer{
			ID:        item.ID,
			Firstname: item.Firstname,
			Lastname:  item.Lastname,
			Email:     item.Email,
			Address:   item.Address,
			Phone:     item.Phone,
		}
	}

	response := fiber.Map{
		"records":         allUsers,
		"totalRecord":     dbCount,
		"number_pages":    pagesNumber,
		"record_per_page": pageSize,
		"links":           links,
	}

	output := fiber.Map{
		"status":  false,
		"message": "no data found",
		"data":    response,
	}
	return c.Status(200).JSON(output)
}

func Paginate(c *fiber.Ctx, pageSize int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
