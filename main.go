package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/Shopify/sarama"
)

var Version string
var Build string
var Producer sarama.SyncProducer

func getenv(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func handler(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte(getenv("GITHUB_SECRET","")))
	if err != nil {
		http.Error(w, "Unauthorized", 401)
	} else {
		if getenv("PAYLOAD_DEBUG", "") == "true" {
			log.Println(string(payload))
		}
		_, _, err := Producer.SendMessage(&sarama.ProducerMessage{
			Topic: getenv("KAFKA_TOPIC","github-event-stream"),
			Value: sarama.StringEncoder(string(payload)),
		})
		if err != nil {
			log.Println(err)
			http.Error(w, "error handling message", 500)
		}
		w.WriteHeader(http.StatusAccepted)
	}
}

func createProducer() sarama.SyncProducer {
	if getenv("KAFKA_DEBUG", "") == "true" {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}
	brokerList := strings.Split(getenv("KAFKA_BROKERS","localhost:9092"), ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatal(err)
	}
	return producer
}

func main() {
	Producer = createProducer()
	mux := http.NewServeMux()
 	mux.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", getenv("PORT","8000")), mux)
}
