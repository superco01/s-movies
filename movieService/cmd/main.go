package main

import (
	"log"
	"net"
	"os"
	"search/movieService/constant"
	"search/movieService/core"
	grpcProto "search/movieService/grpc"
	controller "search/movieService/grpc/controller"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	service := core.MovieSearchService{}
	serviceController := controller.NewMovieSearchController(service)

	grpcServer := grpc.NewServer()
	grpcProto.RegisterMovieSearchServer(grpcServer, serviceController)
	reflection.Register(grpcServer)

	con, err := net.Listen(constant.NETWORK, os.Getenv(constant.KEY_ADDRESS))
	if err != nil {
		log.Panic(err)
	}
	log.Println("grpc application up and running on port" + os.Getenv(constant.KEY_ADDRESS))

	err = grpcServer.Serve(con)
	if err != nil {
		log.Panic(err)
	}
}
