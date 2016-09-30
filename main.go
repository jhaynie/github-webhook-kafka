package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"net/http"
	"encoding/json"

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

func stringify(v interface{}) string {
	buf, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(buf)
}

func handler(w http.ResponseWriter, r *http.Request) {
	payload, err := github.ValidatePayload(r, []byte(getenv("GWK_GITHUB_SECRET","")))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	} else {
		if getenv("GWK_PAYLOAD_DEBUG", "") == "true" {
			log.Println(string(payload))
		}
		name := github.WebHookType(r)
		event, err := github.ParseWebHook(name, payload)
		if err != nil {
			http.Error(w, "Invalid Event", http.StatusInternalServerError)
		} else {
			_, _, err := Producer.SendMessage(&sarama.ProducerMessage{
				Topic: fmt.Sprintf("%s%s", getenv("GWK_KAFKA_TOPIC_PREFIX",""), name),
				Value: sarama.StringEncoder(stringify(event)),
			})
			if err != nil {
				log.Println(err)
				http.Error(w, "error handling message", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func createProducer() sarama.SyncProducer {
	if getenv("GWK_KAFKA_DEBUG", "") == "true" {
		sarama.Logger = log.New(os.Stdout, "[github-webhook-kafka] ", log.LstdFlags)
	}
	brokerList := strings.Split(getenv("GWK_KAFKA_BROKERS","localhost:9092"), ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 100
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
	http.ListenAndServe(fmt.Sprintf(":%s", getenv("GWK_PORT","8000")), mux)
}
