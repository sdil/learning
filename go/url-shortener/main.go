package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	h "github.com/sdil/learning/go/url-shortener/api"
	mr "github.com/sdil/learning/go/url-shortener/repository/mongo"

	"github.com/sdil/learning/go/url-shortener/shortener"
)

func httpPort() string {
	port := "8081"
	return fmt.Sprintf(":%s", port)
}

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

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening to port :8000")
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}
