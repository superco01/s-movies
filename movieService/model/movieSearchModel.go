package model

type MovieRequest struct {
	MovieName string
	Page      int
}

type MovieResponse struct {
	Response     []Movie `json:"Search"`
	ErrorMessage string  `json:"Error"`
}

type Response struct {
	Search       []Movie `json:"Search"`
	TotalResults string  `json:"totalResults"`
	Response     string  `json:"Response"`
	Error        string  `json:"Error"`
}

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}
