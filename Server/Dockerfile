FROM golang:1.15-alpine

WORKDIR /go/src/github.com/Ossamoon/HealthTalk/Server
COPY . .

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"]