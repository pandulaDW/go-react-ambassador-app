package database

import (
	"fmt"
	"github.com/pandulaDW/go-react-ambassador-app/src/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
		return
	}
}

// Connect will connect to the mysql database
func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error connecting to the db: ", err)
		os.Exit(1)
	}
}

// AutoMigrate will migrate the tables
func AutoMigrate() {
	err := DB.AutoMigrate(models.User{}, models.Product{})
	if err != nil {
		fmt.Println("error migrating the tables: ", err)
		os.Exit(1)
	}
}
