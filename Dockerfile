FROM golang:1.25

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o src/main .

EXPOSE 9300

CMD ["./main"]
