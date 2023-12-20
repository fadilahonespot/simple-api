package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fadilahonespot/simple-api/repository"
	"github.com/fadilahonespot/simple-api/server/handler"
	"github.com/fadilahonespot/simple-api/server/router"
	"github.com/fadilahonespot/simple-api/usecase"
	"github.com/fadilahonespot/simple-api/utils/database"
	"github.com/fadilahonespot/simple-api/utils/logger"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load Env
	godotenv.Load()
	
	// Setup database
	db := database.InitDB()

	// Setup logger
	logger.NewLogger()

	// Setup repository
	productRepo := repository.NewProductRepository(db)

	// Setup usecase
	productUsecase := usecase.NewProductRepository(productRepo)

	// Set handler
	productHandler := handler.NewProductHandler(productUsecase)

	// Set Router
	e := echo.New()
	router := router.DefaultRouter{
		ProductHandler: &productHandler,
	}
	router.NewRouter(e).Validate()
	err := e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}