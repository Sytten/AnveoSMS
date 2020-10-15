APP=anveosms

.PHONY: build
build: clean
	go build -o ${APP} main.go

.PHONY: run
run:
	go run -race main.go

.PHONY: clean
clean:
	go clean
