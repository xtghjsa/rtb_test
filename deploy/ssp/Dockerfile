FROM golang:1.23.6-alpine3.21 AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ssp ./main.go

FROM alpine:3.21
WORKDIR /
COPY --from=builder /app/ssp .
COPY .env .
EXPOSE 6060
CMD ["./ssp", "startssp"]