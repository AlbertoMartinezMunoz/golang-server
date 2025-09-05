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
