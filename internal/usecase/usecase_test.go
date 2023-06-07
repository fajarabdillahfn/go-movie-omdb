package usecase

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/fajarabdillahfn/go-movie-omdb/internal/model"
)

func TestSearch(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	apiUrl := "http://www.omdbapi.com/"

	type fields struct {
		apiKey string
		apiUrl string
	}
	type args struct {
		ctx   context.Context
		param *model.SearchParameter
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData struct {
			title      string
			errMessage string
		}
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{apiKey: apiKey, apiUrl: apiUrl},
			args: args{ctx: context.Background(), param: &model.SearchParameter{
				SearchKeyword: "Batman",
			}},
			wantData: struct {
				title      string
				errMessage string
			}{
				"Batman",
				"",
			},
			wantErr: false,
		},
		{
			name:   "no data",
			fields: fields{apiKey: apiKey, apiUrl: apiUrl},
			args: args{ctx: context.Background(), param: &model.SearchParameter{
				SearchKeyword: "aklsdghasdopciajs",
			}},
			wantData: struct {
				title      string
				errMessage string
			}{
				"",
				"Movie not found!",
			},
			wantErr: true,
		},
		{
			name:   "invalid api key",
			fields: fields{apiKey: "", apiUrl: apiUrl},
			args: args{ctx: context.Background(), param: &model.SearchParameter{
				SearchKeyword: "Batman",
			}},
			wantData: struct {
				title      string
				errMessage string
			}{
				"",
				"No API key provided.",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := movieUseCase{
				apiKey: tt.fields.apiKey,
				url:    tt.fields.apiUrl,
			}
			gotData, _ := u.Search(tt.args.ctx, tt.args.param)

			if !tt.wantErr {
				for _, data := range gotData.Search {
					if !strings.Contains(data.Title, tt.wantData.title) {
						t.Errorf("Search error = Title should contains: %v, but got: %v", tt.wantData.title, data.Title)
						return
					}
				}
			} else {
				if gotData.Error != tt.wantData.errMessage {
					t.Errorf("Search error = %v, wantErr %v", gotData.Error, tt.wantErr)
					return
				}
			}

		})
	}
}
