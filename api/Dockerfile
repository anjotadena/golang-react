FROM golang:1.14.2-alpine3.11 as dev
WORKDIR /app
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY . .
CMD modd

# FROM golang:1.14.2-alpine3.11 as build
# WORKDIR /app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .