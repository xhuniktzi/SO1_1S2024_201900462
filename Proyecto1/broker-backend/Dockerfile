FROM golang:1.20-alpine AS build

WORKDIR /usr/src/app

COPY . .

RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]
