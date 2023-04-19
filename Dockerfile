FROM golang:latest

WORKDIR /transaction

COPY go.mod go.sum ./

COPY . .

RUN go mod download


EXPOSE 8000

CMD ["go", "run", "./cmd/main.go"]