FROM golang:1.26.3

WORKDIR /app

COPY . .

RUN go build -o main ./cmd

CMD ["./main"]