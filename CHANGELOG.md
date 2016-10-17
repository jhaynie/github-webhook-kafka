# Changelog

## 1.0.0 - 09/28/2016

- initial release

## 1.0.1 - 09/29/2016

- parse incoming event payload and place the event into the queue instead of the entire payload
- changed `KAFKA_TOPIC` to `KAFKA_TOPIC_PREFIX` and prepend this string to the name of the event to form topic name
- switched to [landoop/fast-data-dev](https://hub.docker.com/r/landoop/fast-data-dev/) for dev container
- renamed debug log prefix from `sarama` to `github-webhook-kafka`
- renamed all environment variables with `GWK_` prefix

## 1.0.2

- set `config.Producer.MaxMessageBytes = 10000000`
- set `config.Producer.Compression = CompressionGZIP`
- in the dev container, set `message.max.bytes=10000000` for kafka broker for larger message bodies by default
- revert to use full event instead of event payload from previous release so additional metadata can be delivered
- changed format of kafka value to be JSON with string payload of body
- ability to specify an encrypted context in the url path which will be automatically decrypted with the provided key and provided in the kafka event payload
- provide an ability to generate a key from command line by passing two arguments (key as 16 characters and text to encrypt)
- switch to emit protobuf encoded messages using [go-github-protobuf](https://github.com/jhaynie/go-github-protobuf)
