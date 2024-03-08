package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"vk2telegram/internal/http"
	"vk2telegram/internal/services"
	"vk2telegram/internal/storage"
)

func main() {
	InitEnv()
	db := InitDB()
	defer db.Close()
	vk := services.NewVKAPI(os.Getenv("VK_API_TOKEN"), os.Getenv("VK_API_URL"), os.Getenv("VK_API_VERSION"))
	telegram := services.NewTelegramAPI(os.Getenv("TELEGRAM_API_TOKEN"), vk)

	go func() {
		telegram.GetUpdates(db)
	}()

	h := new(http.Handler)
	srv := new(http.Server)
	err := srv.Run(os.Getenv("PORT"), h.InitRoutes())
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InitDB() *storage.DatabaseStorage {
	db, err := storage.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}
	//defer db.Close()

	return db
}

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
