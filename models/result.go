package models

type Result struct {
	StatusCode string
	StatusMsg  string
	Error      error
	Body       interface{}
}

type ResultGet struct {
	StatusCode string
	StatusMsg  string
	Error      error
	Body       Paginate
}

type Paginate struct {
	Count  int
	Total  int
	Prices []PriceHistory
}
