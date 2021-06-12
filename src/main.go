package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rayzalzero/go-sukha/src/config"
	middleware "github.com/rayzalzero/go-sukha/src/middlewares"

	_coffeEntity "github.com/rayzalzero/go-sukha/src/api/entity/coffe"
	_coffeHandler "github.com/rayzalzero/go-sukha/src/api/handler/coffe"
	_coffeRepo "github.com/rayzalzero/go-sukha/src/api/repo/coffe"
)

func main() {
	if os.Getenv("ENV") == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error getting env")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "1111"
	}

	timeout := os.Getenv("TIMEOUT")
	if timeout == "" {
		timeout = "2"
	}
	i, _ := strconv.Atoi(timeout)
	timeoutContext := time.Duration(i) * time.Second

	config.Init()
	db := config.GetDB()

	repoDiscuss := _coffeRepo.InitCoffeRepo(db)
	entityDiscuss := _coffeEntity.InitCoffeEntity(repoDiscuss, timeoutContext)

	myfile, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(myfile, os.Stdout)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.RemoveExtraSlash = true

	r.Use(middleware.CORSMiddleware())
	api := r.Group("/")

	_coffeHandler.InitCoffeHandler(api, entityDiscuss)

	r.Run(":" + port)
}
