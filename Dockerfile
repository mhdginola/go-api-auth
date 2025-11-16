FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main src/main.go

EXPOSE 9300

CMD ["./main"]
