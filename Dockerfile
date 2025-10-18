FROM golang:1.24.0
LABEL authors="maxim"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o wp ./cmd

CMD ["./wp"]