FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY order order

# gRPC client services
COPY product product

RUN go build -o order_service ./order/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/order_service /src/app
CMD ["/src/app"]