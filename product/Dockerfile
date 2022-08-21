FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY product product

RUN go build -o product_service ./product/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/product_service /src/app
CMD ["/src/app"]