package main

import (
	"fmt"

	"github.com/hata-okrc/travel_record-app.git/db"
	"github.com/hata-okrc/travel_record-app.git/model"
)

func main() {
	// データベースインスタンスのアドレスを格納
	dbConn := db.NewDB()
	defer fmt.Println("migrate 成功")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Trip{})
}
