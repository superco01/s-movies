# Movie Searcher Services

This project is about and app that can retrieve movie information from various Open API end points and it has 2 services, the mobile service that works as client and the movie service that works as server that listen any grpc request from clients. 
For the first release it only works with OMDB Open API and only implemented minimun viable product, get movie list by movie keyword and page number and get the movie detail using imdbID that has been retrieved from get movie list process.
Communication between services is handled by GRPC and communication to Open API is fully implemented using REST HTTP with JSON as the message type

## Getting Started
### Movie Service
To run this service we can start to build the main file by using this command:
```baseh
cd movieService/cmd/
go build -o ./movieService .
./movieService
```

### Mobile Service
To run this service we can start to build the main file by using previous command in Movie Service or we can just run the code by following this command:
```bash
cd mobileService/cmd/
go run main.go
```
