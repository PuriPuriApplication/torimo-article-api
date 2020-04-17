FROM golang:1.13.7-alpine3.11

WORKDIR /torimo-article-api

ENV GO111MODULE=on

COPY . .

# TODO: main.go のあるディレクトリでビルドを実行しないとうまくいかない (wire)
RUN go get github.com/google/wire/cmd/wire \
  && wire src/wire.go

CMD ["go", "test", "-v", "./..."]
