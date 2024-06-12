FROM golang:1.22.1

WORKDIR /app/gamer_shop

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-gs .

EXPOSE 8080

CMD ["./api-gs"]
