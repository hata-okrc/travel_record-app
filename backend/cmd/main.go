package main

import (
	"github.com/hata-okrc/travel_record-app.git/controller"
	"github.com/hata-okrc/travel_record-app.git/db"
	"github.com/hata-okrc/travel_record-app.git/repository"
	"github.com/hata-okrc/travel_record-app.git/router"
	"github.com/hata-okrc/travel_record-app.git/usecase"
)


func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))

}