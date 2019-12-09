# version=1.0

all: build

rundev:
	DOTENV=true go run main.go

build:
	go build -o lapis main.go

clean:
	rm lapis

run:
	./lapis