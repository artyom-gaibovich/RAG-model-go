package main

import (
	"github.com/APSN4/RAG-model-go/app/router"
	"github.com/APSN4/RAG-model-go/config"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
