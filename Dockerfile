##
## Build
##

FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

ENV CGO_ENABLED 0

RUN go build -mod=vendor -o /snappfood_test main.go

##
## Deploy
##

FROM alpine:3.14

WORKDIR /

COPY --from=build /snappfood_test /snappfood_test
COPY . /data

EXPOSE 8080

CMD ["/snappfood_test"]