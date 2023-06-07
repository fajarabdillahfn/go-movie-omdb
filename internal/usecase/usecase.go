package usecase

import (
	"context"

	"github.com/fajarabdillahfn/go-movie-omdb/internal/model"
)

type MovieUseCase interface {
	Search(ctx context.Context, param *model.SearchParameter) (*model.SearchResult, error)
	GetByID(ctx context.Context, id string) (*model.MovieDetail, error)
}

type movieUseCase struct {
	apiKey string
}

func (u *movieUseCase) Search(ctx context.Context, param *model.SearchParameter) (*model.SearchResult, error) {
	return nil, nil
}

func (u *movieUseCase) GetByID(ctx context.Context, id string) (*model.MovieDetail, error) {
	return nil, nil
}