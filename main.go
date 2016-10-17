package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"time"
	"errors"
	"bytes"
	"strings"
	"os/exec"
	"os/signal"
	"net/http"
	"crypto/rand"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"

	"github.com/jhaynie/go-github-protobuf/github"
	"github.com/golang/protobuf/proto"
	"github.com/satori/go.uuid"
	"github.com/Shopify/sarama"
)

var Version string
var Build string
var Key string
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


func decrypt(text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}
	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func encrypt(text []byte) (string, error) {
	block, err := aes.NewCipher([]byte(Key))
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	var context string
	if Key != "" {
		// if we have a configured Key and we are passed in a context,
		// we are going to decrypt it with the provided Key
		buf, err := base64.StdEncoding.DecodeString(r.URL.Path[1:])
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		dec, err := decrypt(buf)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		context = dec
	}
	payload, err := github.ValidatePayload(r, []byte(getenv("GWK_GITHUB_SECRET","")))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	} else {
		name := github.WebHookType(r)
		log.Println(string(payload))
		event, err := github.ParseWebHook(name, payload)
		if err != nil {
			log.Println(err)
			http.Error(w, "error handling message", http.StatusInternalServerError)
		} else {
			protobuf, err := proto.Marshal(event.(proto.Message))
			if err != nil {
				log.Println(err)
				http.Error(w, "error handling message", http.StatusInternalServerError)
			} else {
				message := Message{
					Version: 1,
					Timestamp: time.Now().UTC().Format(time.RFC3339),
					Id: uuid.NewV4().String(),
					Event: name,
					Payload: protobuf,
					Context: context,
				}
				log.Println(stringify(message))
				value, err := proto.Marshal(&message)
				if err != nil {
					log.Println(err)
					http.Error(w, "error handling message", http.StatusInternalServerError)
				} else {
					topic := fmt.Sprintf("%s%s", getenv("GWK_KAFKA_TOPIC_PREFIX",""), name)
					if getenv("GWK_PAYLOAD_DEBUG", "") == "true" {
						log.Println(fmt.Sprintf("sending to %s", topic))
					}
					_, _, err := Producer.SendMessage(&sarama.ProducerMessage{
						Topic: topic,
						Value: sarama.ByteEncoder(value),
					})
					if err != nil {
						log.Println(err)
						http.Error(w, "error handling message", http.StatusInternalServerError)
					}
					w.WriteHeader(http.StatusAccepted)
				}
			}
		}
	}
}

func createProducer() sarama.SyncProducer {
	if getenv("GWK_KAFKA_DEBUG", "") == "true" {
		sarama.Logger = log.New(os.Stdout, "[github-webhook-kafka] ", log.LstdFlags)
	}
	brokerList := strings.Split(getenv("GWK_KAFKA_BROKERS","localhost:9092"), ",")
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP
	config.Producer.MaxMessageBytes = 10000000
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 100
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatal(err)
	}
	return producer
}

func main() {
	if len(os.Args) > 2 {
		Key = os.Args[1]
		buf, err := encrypt([]byte(os.Args[2]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf)
		os.Exit(0)
	}
	Key = getenv("GWK_AES_KEY","")
	Producer = createProducer()
	mux := http.NewServeMux()
 	mux.HandleFunc("/", handler)
	listen := fmt.Sprintf(":%s", getenv("GWK_PORT","8000"))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func(){
		for range c {
			log.Println("[github-webhook-kafka] SIGINT, exiting...")
			Producer.Close()
			os.Exit(0)
		}
	}()
	go func(){
		cmd := exec.Command("lt", "--port", "8000")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 10; i++ {
			s := strings.TrimSpace(out.String())
			if s == "" {
				time.Sleep(1 * time.Second)
				continue
			}
	    	fmt.Printf("[github-webhook-kafka] URL: \033[31m%s\033[m\n", strings.Replace(s, "your url is:", "", -1))
			break
		}
	}()
	log.Printf("[github-webhook-kafka] listening on port %s\n", listen[1:])
	err := http.ListenAndServe(listen, mux)
	if err != nil {
		log.Fatal(err)
	}
}
