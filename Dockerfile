FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server_generation

FROM alpine

WORKDIR /build

COPY --from=builder /build/server_generation /build/server_generation
COPY --from=builder /build/.env /build/.env

CMD ["/build/server_generation"]