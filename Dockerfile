FROM alpine:3.4

MAINTAINER Jeff Haynie <jhaynie@gmail.com>

EXPOSE 8000

RUN mkdir -p /app

COPY build/alpine/github-webhook-kafka /app

WORKDIR /app

CMD ["/app/github-webhook-kafka"]
