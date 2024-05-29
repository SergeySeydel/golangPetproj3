BINARY_NAME=goApp

all: build

build:

	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

.PHONY: all build run clean
