FROM golang:1.15-alpine

WORKDIR /go/src/github.com/Ossamoon/HealthTalk/Server
COPY . .

RUN apk update && \
    apk add --no-cache git && \
    go get github.com/pilu/fresh

RUN apk add --no-cache openssl

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

EXPOSE 8080

CMD ["fresh"]