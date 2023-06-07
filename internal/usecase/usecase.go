package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/fajarabdillahfn/go-movie-omdb/internal/model"
)

type MovieUseCase interface {
	Search(ctx context.Context, param *model.SearchParameter) (*model.SearchResult, *model.ErrorResponse)
	GetByID(ctx context.Context, id string) (*model.MovieDetail, *model.ErrorResponse)
}

type movieUseCase struct {
	apiKey string
	url    string
}

func (u *movieUseCase) Search(ctx context.Context, param *model.SearchParameter) (*model.SearchResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	searchResult := model.SearchResult{}

	releaseYearStr := strconv.Itoa(int(param.ReleaseYear))
	pageStr := strconv.Itoa(int(param.Page))

	if releaseYearStr == "0" {
		releaseYearStr = ""
	}
	if pageStr == "0" {
		pageStr = ""
	}

	params := url.Values{}
	params.Add("apiKey", u.apiKey)
	params.Add("s", param.SearchKeyword)
	params.Add("type", param.DataType)
	params.Add("y", releaseYearStr)
	params.Add("r", param.DataFormat)
	params.Add("page", pageStr)
	params.Add("callback", param.Callback)

	urlStr := fmt.Sprintf("%s?%s", u.url, params.Encode())

	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&searchResult)
	if err != nil {
		return nil, err
	}

	return &searchResult, nil
}

func (u *movieUseCase) GetByID(ctx context.Context, id string) (*model.MovieDetail, *model.ErrorResponse) {
	return nil, nil
}
