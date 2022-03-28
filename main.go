package main

import (
	"fmt"
	"log"
	"myGram/packages"

	"myGram/packages/database"
	"myGram/src/controllers"
	"myGram/src/repositories"
	"myGram/src/routers"
	"myGram/src/usecases"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.LoadDB()
}

func main() {
	db := database.GetDB()
	repository := repositories.NewRepo(db)
	usecase := usecases.NewUC(repository)
	ctrl := controllers.NewCtrl(usecase)
	route := routers.NewRouter(ctrl).RouterGroup()
	log.Printf("[SERVER] starting at port : %v", os.Getenv("SERVER_PORT"))
	log.Fatalln(
		http.ListenAndServe(
			fmt.Sprintf(":%v", os.Getenv("SERVER_PORT")),
			packages.AllowCORS(route),
		),
	)
}
