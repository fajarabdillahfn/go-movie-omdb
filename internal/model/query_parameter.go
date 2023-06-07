package model

type SearchParameter struct {
	SearchKeyword string
	DataType      string
	ReleaseYear   uint
	DataFormat    string
	Page          uint
	Callback      string
}
