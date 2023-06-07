package http

import (
	"net/http"
	"strconv"

	"github.com/fajarabdillahfn/go-movie-omdb/internal/model"
	"github.com/fajarabdillahfn/go-movie-omdb/internal/usecase"
	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	usecase usecase.MovieUseCase
}

func NewMovieHandler(usecase usecase.MovieUseCase) *MovieHandler {
	return &MovieHandler{
		usecase: usecase,
	}
}

func (h *MovieHandler) Route(r *gin.RouterGroup) {
	r.GET("/search", h.Search)
	r.GET("/detail/:id", h.GetByID)
}

func (h *MovieHandler) Search(c *gin.Context) {
	releaseYearInt, _ := strconv.Atoi(c.Query("y"))
	pageInt, _ := strconv.Atoi(c.Query("page"))

	param := model.SearchParameter{
		SearchKeyword: c.Query("s"),
		DataType:      c.Query("type"),
		ReleaseYear:   uint(releaseYearInt),
		DataFormat:    c.Query("r"),
		Page:          uint(pageInt),
		Callback:      c.Query("callback"),
	}

	result, err := h.usecase.Search(c, &param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *MovieHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.usecase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result)
}
