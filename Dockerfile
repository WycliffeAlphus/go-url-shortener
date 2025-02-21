
FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Make sure the database directory exists and is writable
RUN mkdir -p /app/database && chmod 777 /app/database

EXPOSE 8080

CMD ["./main"]