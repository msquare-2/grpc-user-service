# Dockerfile

FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o grpc-user-service ./main.go

EXPOSE 50051

CMD ["./grpc-user-service"]
