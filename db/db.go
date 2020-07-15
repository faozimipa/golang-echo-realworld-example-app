package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// "github.com/joho/godotenv"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/faozimipa/golang-echo-realworld-example-app/model"

)

//New db func 
func New() *gorm.DB {
	// DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
		
	// db, err := gorm.Open("postgres", DBURL)
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=rw password=secret  sslmode=disable")

	if err != nil {
		fmt.Println("Database error: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

//TestDB func 
func TestDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=rw password=secret  sslmode=disable")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

//DropTestDB func
func DropTestDB() error {
	if err := os.Remove("./../realworld_test.db"); err != nil {
		return err
	}
	return nil
}


//AutoMigrate func 
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
		&model.Article{},
		&model.Comment{},
		&model.Tag{},
	)
}
