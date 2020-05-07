FROM golang:1.13.7-alpine3.11 as build

WORKDIR /torimo-article-api

ENV GO111MODULE=on

RUN apk add --no-cache git \
  && GO111MODULE=off go get github.com/oxequa/realize

COPY . .

# TODO: main.go のあるディレクトリでビルドを実行しないとうまくいかない (wire)
RUN go get github.com/google/wire/cmd/wire \
  && wire src/wire.go \
  && cd src \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../main \
  && cd ..


FROM alpine

COPY --from=build /torimo-article-api/main .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go ./main

CMD ["./main"]
