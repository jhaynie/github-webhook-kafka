FROM alpine:3.4

MAINTAINER Jeff Haynie <jhaynie@gmail.com>

ENV GWK_VERSION 1.0.0
ENV GWK_URL https://github.com/jhaynie/github-webhook-kafka/releases/download/"$GWK_VERSION"/github-webhook-kafka-alpine-"$GWK_VERSION"

EXPOSE 8000

RUN mkdir -p /app

RUN apk update && \
	apk add ca-certificates wget && \
	update-ca-certificates && \
	wget -q $GWK_URL -O /app/github-webhook-kafka && \
	chmod +x /app/github-webhook-kafka

WORKDIR /app

CMD ["/app/github-webhook-kafka"]
