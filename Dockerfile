FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o ./build/executable/app ./cmd/main.go

CMD ["./build/executable/app"]