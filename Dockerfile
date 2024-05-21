FROM golang:1.20

WORKDIR /app
#
COPY go.mod go.sum ./
#

RUN go mod download

COPY . .

RUN go mod vendor

RUN mkdir /go/logs && CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -o main


# Expose 8080
EXPOSE 8080

CMD ["go", "run", "main.go"]