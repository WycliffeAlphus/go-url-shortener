FROM golang:1.22-alpine

RUN apk add --no-cache sqlite-dev build-base

WORKDIR /app

COPY . .

RUN go mod download


RUN CGO_ENABLED=1 go build -o main .

RUN mkdir -p /app/database && chmod 777 /app/database

EXPOSE 8080

CMD ["./main"]