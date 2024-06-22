FROM golang:1.22.1

WORKDIR /app/gamer-shop/back-end

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o api-gs .

EXPOSE 8080

CMD ["./api-gs"]
