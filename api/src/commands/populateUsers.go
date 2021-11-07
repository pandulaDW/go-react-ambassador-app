package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/models"
)

func main() {
	database.Connect()
	database.DB.Exec("DELETE FROM users")
	database.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

	users := make([]models.User, 0, 30)
	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}
		ambassador.SetPassword([]byte("1234"))
		users = append(users, ambassador)
	}

	tx := database.DB.Create(&users)

	fmt.Printf("Inserted %d users", tx.RowsAffected)
	database.CloseConnection()
}
