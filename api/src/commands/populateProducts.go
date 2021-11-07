package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/pandulaDW/go-react-ambassador-app/src/database"
	"github.com/pandulaDW/go-react-ambassador-app/src/models"
	"math"
)

func main() {
	database.Connect()
	database.DB.Exec("DELETE FROM products")
	database.DB.Exec("ALTER TABLE products AUTO_INCREMENT = 1")

	products := make([]models.Product, 0, 30)
	for i := 0; i < 30; i++ {
		product := models.Product{
			Title:       faker.Word(),
			Description: faker.Sentence(),
			Image:       faker.URL(),
			Price:       math.Abs(faker.Latitude()),
		}
		products = append(products, product)
	}

	tx := database.DB.Create(&products)
	fmt.Printf("Inserted %d products", tx.RowsAffected)

	database.CloseConnection()
}
