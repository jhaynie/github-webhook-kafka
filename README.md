# Github Webhook -> Kafka

This is a simple Github web hook server which streams incoming events into a Kafka queue.  Each incoming web hook event will be validated with the Github secret and any incoming requests not matching the validation will be rejected with a HTTP status code `401 Unauthorized`.

## Configuration

The following environment variables need to be set:

- `GITHUB_SECRET` - this is the secret used when creating / registering the web hook.  Defaults to empty string
- `KAFKA_TOPIC` - the kafka topic name.  defaults to `github-event-stream`
- `KAFKA_BROKERS` - a comma separated list of Kafka servers. Defaults to `localhost:9092`
- `PORT` - webserver HTTP port. defaults to `8000`

Optionally, you can control debugging with:

- `PAYLOAD_DEBUG` - dump each incoming JSON event payload to the console when `true`
- `KAFKA_DEBUG` - log kafka details when `true`

## Running

You can easily run using Docker.

If you want to use a pre-built container ([`jhaynie/github-webhook-kafka`](https://hub.docker.com/r/jhaynie/github-webhook-kafka/)), you could use something like:

`docker run -d -p 8000:8000 -e GITHUB_SECRET=hello -e KAFKA_BROKERS=192.168.1.1:9092 jhaynie/github-webhook-kafka`

Make sure you use the correct `GITHUB_SECRET` and `KAFKA_BROKERS` values.  This will run the container on port 8000 by default.

For development, you can use the all-in-one container which contains Kafka, Zookeeper and this software:

`docker run -d -p 8000:8000 -p 2181:2181 -p 9092:9092 -e GITHUB_SECRET=hello -e ADVERTISED_HOST=192.168.1.1 jhaynie/github-webhook-kafka-dev`

Make sure you use the correct `GITHUB_SECRET`, `ADVERTISED_HOST` values.


## License

Copyright (c) 2016 by Jeff Haynie

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
