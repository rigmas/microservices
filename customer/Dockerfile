FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY customer customer

RUN go build -o customer_service ./customer/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/customer_service /src/app
CMD ["/src/app"]