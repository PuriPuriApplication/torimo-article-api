build:
	go build -o bin/torimo-article-api src/*.go

clean:
	rm -rf ./bin

run:
    go run main.go