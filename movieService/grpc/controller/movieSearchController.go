package controller

import (
	"context"
	"search/movieService/core"
	grpcProto "search/movieService/grpc"
	"search/movieService/model"
)

type MovieSearchController struct {
	movieSearch core.Service
}

func NewMovieSearchController(movieSearch core.Service) grpcProto.MovieSearchServer {
	return &MovieSearchController{
		movieSearch: movieSearch,
	}
}

func (controller *MovieSearchController) GetMovies(ctx context.Context, request *grpcProto.MovieRequest) (response *grpcProto.MovieResponse, err error) {
	movies, err := controller.movieSearch.GetMovies(request.Keyword, int(request.Page))
	if err != nil {
		return
	}

	// var moviesArr []*grpcProto.Movie
	response = &grpcProto.MovieResponse{}
	for _, v := range movies {
		response.Movies = append(response.Movies, marshalMovie(&v))
	}

	// response.Movies = moviesArr

	return
}

func marshalMovie(m *model.Movie) *grpcProto.Movie {
	return &grpcProto.Movie{
		Title:  m.Title,
		Year:   m.Year,
		Type:   m.Type,
		ImdbID: m.ImdbID,
		Poster: m.Poster,
	}
}
