FROM golang:1.25-alpine

WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
