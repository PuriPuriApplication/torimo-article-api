FROM golang:1.13.7-alpine3.11

WORKDIR /app

ENV GO111MODULE=on

COPY . .

CMD ["go", "run", "src/main.go"]
