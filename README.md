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

