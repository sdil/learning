package main

import (
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"

	h "github.com/sdil/learning/go/url-shortener/api"
	mr "github.com/sdil/learning/go/url-shortener/repository/mongo"
	"github.com/sdil/learning/go/url-shortener/shortener"
)

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "mongo":
		log.Println("Using MongoDB")
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	log.Println("No database")
	panic("No database selected")
	return nil
}

func main() {
	repo := chooseRepo()
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	e := echo.New()
	e.GET("/:code", handler.Get)
	e.POST("/", handler.Post)

	e.Logger.Fatal(e.Start(":8081"))
}
