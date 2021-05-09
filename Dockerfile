FROM golang:1.14.0-stretch

RUN mkdir /app

ADD ./app /app

WORKDIR /app

RUN go mod download

RUN go build -o server .

CMD ["/app/server"]