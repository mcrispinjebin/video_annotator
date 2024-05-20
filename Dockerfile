FROM golang:1.21

WORKDIR /app
#
COPY go.mod go.sum ./
#

RUN go mod download

COPY . .

RUN mkdir /go/logs && CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -o main -mod vendor


# Expose 8080
EXPOSE 8080

CMD ["go", "run", "main.go"]