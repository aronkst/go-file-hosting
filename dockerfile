FROM golang:1.18-alpine AS build

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN mkdir data
RUN mkdir web

COPY data/*.go data
COPY web/*.go web
COPY main.go .

RUN go build -o go-file-hosting main.go

FROM alpine:3.16

WORKDIR /usr/src/app

RUN mkdir static

COPY --from=build /usr/src/app/go-file-hosting go-file-hosting

CMD [ "./go-file-hosting" ]
