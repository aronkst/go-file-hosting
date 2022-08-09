FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN mkdir data
RUN mkdir web

COPY data/*.go data
COPY web/*.go web
COPY main.go .

RUN mkdir static

RUN go build -o go-file-hosting main.go

CMD [ "./go-file-hosting" ]
