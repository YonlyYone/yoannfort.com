package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)


func main() {
    err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier .env:", err)
		return
	}

    dbHost := os.Getenv("ENV")

    app := fiber.New()

    app.Use(logger.New(logger.Config{
        Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
    }))

    cacheDuration := 10 * time.Second

    if( dbHost != "dev" ) {
        cacheDuration = 0
    } 

    app.Static("/", "./website/build", fiber.Static{
        Browse:        true,
        Index: "index.html",
        CacheDuration: cacheDuration,
        MaxAge:        3600,
    }) 

    log.Fatal(app.Listen(":3000"))
}