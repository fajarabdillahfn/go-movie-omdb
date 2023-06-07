package usecase

import (
	"context"

	"github.com/fajarabdillahfn/go-movie-omdb/internal/model"
)

type MovieUseCase interface {
	Search(ctx context.Context, param *model.SearchParameter) (*model.SearchResult, *model.ErrorResponse)
	GetByID(ctx context.Context, id string) (*model.MovieDetail, *model.ErrorResponse)
}

type movieUseCase struct {
	apiKey string
	url string
}

func (u *movieUseCase) Search(ctx context.Context, param *model.SearchParameter) (*model.SearchResult, *model.ErrorResponse) {
	return nil, nil
}

func (u *movieUseCase) GetByID(ctx context.Context, id string) (*model.MovieDetail, *model.ErrorResponse) {
	return nil, nil
}
