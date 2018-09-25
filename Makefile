# version=1.0

all: setup build

setup:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

build:
	go build -o lapis main.go

clean:
	rm lapis

run:
	./lapis