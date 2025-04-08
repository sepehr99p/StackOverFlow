FROM golang:1.24.1

WORKDIR /app

RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN /go/bin/swag init --output ./docs

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
