serve: 
	go run main.go 

build: 
	go build -o ./build/server main.go 

test: 
	go test ./... -v -gcflags=-l
