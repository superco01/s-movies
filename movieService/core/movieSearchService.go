package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"search/movieService/model"
	"search/movieService/util"
)

type Service interface {
	GetMovies(keyword string, page int) ([]model.Movie, error)
}

type MovieSearchService struct{}

func NewService() Service {
	return MovieSearchService{}
}

func (MovieSearchService) GetMovies(keyword string, page int) ([]model.Movie, error) {
	var endPoint string = util.NewUtil().ComposeOMDBEndPoint(keyword, page)

	response, err := http.Get(endPoint)
	if err != nil {
		fmt.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Search)
	return responseObject.Search, err
}
