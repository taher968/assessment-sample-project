


build:
	go build -o main cmd/main.go


run: 
	go run cmd/main.go

test: 
	go test ./...

clean: 
	rm -f main