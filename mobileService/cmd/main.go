package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"search/mobileService/constant"
	"search/movieService/core"
	msgrpc "search/movieService/grpc"
	"search/movieService/model"

	"google.golang.org/grpc"
)

type grpcService struct {
	grpcClient msgrpc.MovieSearchClient
}

var pageGlobal int
var keywordGlobal string
var moviesGlobal []model.Movie

func NewGRPCService(connectionString string) (core.Service, error) {
	conn, err := grpc.Dial(connectionString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &grpcService{grpcClient: msgrpc.NewMovieSearchClient(conn)}, nil
}

func (s *grpcService) GetMovies(keyword string, page int) (result []model.Movie, err error) {
	req := &msgrpc.MovieRequest{Keyword: keyword, Page: int64(page)}

	ctx, cancelFunc := context.WithTimeout(context.Background(), constant.DEFAULT_REQUEST_TIMEOUT)
	defer cancelFunc()

	resp, err := s.grpcClient.GetMovies(ctx, req)
	if err != nil {
		return
	}

	for _, v := range resp.Movies {
		movie := Unmarshal(v)
		result = append(result, movie)
	}

	if result != nil {
		keywordGlobal = keyword
		pageGlobal = page
		moviesGlobal = result
	}

	return
}

func GetMovieDetail(ImdbID string, grpcService core.Service) (result model.Movie, err error) {

	if moviesGlobal == nil {
		grpcService.GetMovies(keywordGlobal, pageGlobal)
	}

	for _, v := range moviesGlobal {
		if v.ImdbID == ImdbID {
			result.Title = v.Title
			result.Year = v.Year
			result.Type = v.Type
			result.ImdbID = v.ImdbID
			result.Poster = v.Poster
		}
	}

	return
}

func Unmarshal(grpcMovie *msgrpc.Movie) (result model.Movie) {
	result.Title = grpcMovie.Title
	result.Year = grpcMovie.Year
	result.Type = grpcMovie.Type
	result.ImdbID = grpcMovie.ImdbID
	result.Poster = grpcMovie.Poster
	return
}

func main() {

	grpcService, err := NewGRPCService(":8000")
	if err != nil {
		log.Println("error instantiating grpcService, error: " + err.Error())
		os.Exit(1)
	}

	movies, errRemote := grpcService.GetMovies("batman", 1)
	if errRemote != nil {
		log.Println("error on calling grpc GetMovies, error: " + errRemote.Error())
		os.Exit(1)
	}
	fmt.Println(movies)
	fmt.Println("--------------------------------")
	movieDetail, err := GetMovieDetail("", grpcService)
	if err != nil {
		log.Println("movie's not found in current page, please call the GetMovies again. error: " + err.Error())
		os.Exit(1)
	}
	fmt.Println(movieDetail)
}
