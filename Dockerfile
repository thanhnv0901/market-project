# https://dev.to/plutov/docker-and-go-modules-3kkn
FROM golang:1.17.2-alpine

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /market_apis

EXPOSE 3500

ENTRYPOINT ["/market_apis"]