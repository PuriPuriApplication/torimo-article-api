FROM golang:1.13.7-alpine3.11

WORKDIR /torimo-article-api

ENV GO111MODULE=on

COPY . .

# TODO: main.go のあるディレクトリでビルドを実行しないとうまくいかない (wire)
RUN cd src \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../main \
  && cd ..

CMD ["./main"]
