package main

import (
	"os"

	handler "github.com/fajarabdillahfn/go-movie-omdb/internal/delivery/http"
	"github.com/fajarabdillahfn/go-movie-omdb/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	apiKey := os.Getenv("API_KEY")
	url := "http://www.omdbapi.com/"

	initializedMovieService(apiKey, url).Route(&r.RouterGroup)

	r.Run(":" + os.Getenv("PORT"))
}

func initializedMovieService(apiKey, url string) *handler.MovieHandler {
	movieUsecase := usecase.NewMovieUseCase(apiKey, url)

	return handler.NewMovieHandler(movieUsecase)
}
