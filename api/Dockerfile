FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/user ./cmd/user-api-server/main.go

EXPOSE 3000

ENTRYPOINT [ "/app/user", "--port", "3000" ]