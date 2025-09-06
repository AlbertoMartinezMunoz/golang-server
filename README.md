# golang-server
Example of a multithread golang server

## Build

To run the program, we use `go run`.
	
```golang
$ go run ./src/main.go
```

To create the binary file we use `go build`.

```golang
GOOS=linux GOARCH=amd64 go build -trimpath ./src/main.go
```

## Test

CURL will be used for testing the web server

To check the server is up a running:

```console
$ curl http://localhost:8080/
```

To check the server runs the correct handler for a request:

```console
$ curl http://localhost:8080/catalogue/the-cure-one --include --header "Content-Type: application/json" --request "GET"
```

## Add Packages

To add new packages, `go get`can be used:

```console
$ go get -u github.com/gosimple/slug
```

## References

* [Go REST Guide. The Standard Library](https://www.jetbrains.com/guide/go/tutorials/rest_api_series/stdlib/)
* [Building a RESTful API with Go: A Step-by-Step Guide](https://medium.com/@briankworld/building-a-restful-api-with-go-a-step-by-step-guide-d17e69f004a7)
* [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
