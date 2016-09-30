# Changelog

## 1.0.0 - 09/28/2016

- initial release

## 1.0.1 - 09/29/2016

- parse incoming event payload and place the event into the queue instead of the entire payload
- changed `KAFKA_TOPIC` to `KAFKA_TOPIC_PREFIX` and prepend this string to the name of the event to form topic name
- switched to [landoop/fast-data-dev](https://hub.docker.com/r/landoop/fast-data-dev/) for dev container
- renamed debug log prefix from `sarama` to `github-webhook-kafka`
- renamed all environment variables with `GWK_` prefix
