package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatalln(err)
		}
	}
	// url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	// url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"),
	//  os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_DB"),  os.Getenv("POSTGRES_PORT"))
	

	// dsn := "host=localhost user=ht password=psht dbname=myapp-db port=5432"

	user := "ht"
	pw := "psht"
	psdb := "myapp-db"
	port := "5432"
	host := "travel_record-app-db-1"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, pw, psdb, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
