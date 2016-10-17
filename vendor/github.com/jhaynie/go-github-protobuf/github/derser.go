package github

// this file was modified heavily from this https://github.com/google/go-github/blob/master/github/activity_events.go
// Portions Copyright 2013 The go-github AUTHORS. All rights reserved.

// we are using our protobuf generated events instead of the ones in their package so that we can marshal into protobuf protocol

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	// sha1Prefix is the prefix used by GitHub before the HMAC hexdigest.
	sha1Prefix = "sha1"
	// sha256Prefix and sha512Prefix are provided for future compatibility.
	sha256Prefix = "sha256"
	sha512Prefix = "sha512"
	// signatureHeader is the GitHub header key used to pass the HMAC hexdigest.
	signatureHeader = "X-Hub-Signature"
	// eventTypeHeader is the Github header key used to pass the event type.
	eventTypeHeader = "X-Github-Event"
)

var (
	// eventTypeMapping maps webhooks types to their corresponding go-github struct types.
	eventTypeMapping = map[string]string{
		"commit_comment":                        "CommitCommentEvent",
		"create":                                "CreateEvent",
		"delete":                                "DeleteEvent",
		"deployment":                            "DeploymentEvent",
		"deployment_status":                     "DeploymentStatusEvent",
		"fork":                                  "ForkEvent",
		"gollum":                                "GollumEvent",
		"integration_installation":              "IntegrationInstallationEvent",
		"integration_installation_repositories": "IntegrationInstallationRepositoriesEvent",
		"issue_comment":                         "IssueCommentEvent",
		"issues":                                "IssuesEvent",
		"member":                                "MemberEvent",
		"membership":                            "MembershipEvent",
		"page_build":                            "PageBuildEvent",
		"public":                                "PublicEvent",
		"pull_request_review_comment":           "PullRequestReviewCommentEvent",
		"pull_request":                          "PullRequestEvent",
		"push":                                  "PushEvent",
		"repository":                            "RepositoryEvent",
		"release":                               "ReleaseEvent",
		"status":                                "StatusEvent",
		"team_add":                              "TeamAddEvent",
		"watch":                                 "WatchEvent",
	}
)

// genMAC generates the HMAC signature for a message provided the secret key
// and hashFunc.
func genMAC(message, key []byte, hashFunc func() hash.Hash) []byte {
	mac := hmac.New(hashFunc, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// checkMAC reports whether messageMAC is a valid HMAC tag for message.
func checkMAC(message, messageMAC, key []byte, hashFunc func() hash.Hash) bool {
	expectedMAC := genMAC(message, key, hashFunc)
	return hmac.Equal(messageMAC, expectedMAC)
}

// messageMAC returns the hex-decoded HMAC tag from the signature and its
// corresponding hash function.
func messageMAC(signature string) ([]byte, func() hash.Hash, error) {
	if signature == "" {
		return nil, nil, errors.New("missing signature")
	}
	sigParts := strings.SplitN(signature, "=", 2)
	if len(sigParts) != 2 {
		return nil, nil, fmt.Errorf("error parsing signature %q", signature)
	}

	var hashFunc func() hash.Hash
	switch sigParts[0] {
	case sha1Prefix:
		hashFunc = sha1.New
	case sha256Prefix:
		hashFunc = sha256.New
	case sha512Prefix:
		hashFunc = sha512.New
	default:
		return nil, nil, fmt.Errorf("unknown hash type prefix: %q", sigParts[0])
	}

	buf, err := hex.DecodeString(sigParts[1])
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding signature %q: %v", signature, err)
	}
	return buf, hashFunc, nil
}

// ValidatePayload validates an incoming GitHub Webhook event request
// and returns the (JSON) payload.
// secretKey is the GitHub Webhook secret message.
//
// Example usage:
//
//     func (s *GitHubEventMonitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//       payload, err := github.ValidatePayload(r, s.webhookSecretKey)
//       if err != nil { ... }
//       // Process payload...
//     }
//
func ValidatePayload(r *http.Request, secretKey []byte) (payload []byte, err error) {
	payload, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	sig := r.Header.Get(signatureHeader)
	if err := validateSignature(sig, payload, secretKey); err != nil {
		return nil, err
	}
	return payload, nil
}

// validateSignature validates the signature for the given payload.
// signature is the GitHub hash signature delivered in the X-Hub-Signature header.
// payload is the JSON payload sent by GitHub Webhooks.
// secretKey is the GitHub Webhook secret message.
//
// GitHub docs: https://developer.github.com/webhooks/securing/#validating-payloads-from-github
func validateSignature(signature string, payload, secretKey []byte) error {
	messageMAC, hashFunc, err := messageMAC(signature)
	if err != nil {
		return err
	}
	if !checkMAC(payload, messageMAC, secretKey, hashFunc) {
		return errors.New("payload signature check failed")
	}
	return nil
}

// WebHookType returns the event type of webhook request r.
func WebHookType(r *http.Request) string {
	return r.Header.Get(eventTypeHeader)
}

// ParseWebHook parses the event payload. For recognized event types, a
// value of the corresponding struct type will be returned (as returned
// by Event.Payload()). An error will be returned for unrecognized event
// types.
//
// Example usage:
//
//     func (s *GitHubEventMonitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//       payload, err := github.ValidatePayload(r, s.webhookSecretKey)
//       if err != nil { ... }
//       event, err := github.ParseWebHook(github.WebHookType(r), payload)
//       if err != nil { ... }
//       switch event := event.(type) {
//       case CommitCommentEvent:
//           processCommitCommentEvent(event)
//       case CreateEvent:
//           processCreateEvent(event)
//       ...
//       }
//     }
//
func ParseWebHook(messageType string, payload []byte) (interface{}, error) {
	eventType, ok := eventTypeMapping[messageType]
	if !ok {
		return nil, fmt.Errorf("unknown X-Github-Event in message: %v", messageType)
	}

	event := Event{
		Type:       &eventType,
		RawPayload: (*json.RawMessage)(&payload),
	}
	return event.Payload(), nil
}

// Event represents a GitHub event.
type Event struct {
	Type       *string          `json:"type,omitempty"`
	Public     *bool            `json:"public"`
	RawPayload *json.RawMessage `json:"payload,omitempty"`
	Repo       *Repository      `json:"repo,omitempty"`
	Actor      *User            `json:"actor,omitempty"`
	Org        *User            `json:"org,omitempty"`
	CreatedAt  *time.Time       `json:"created_at,omitempty"`
	ID         *string          `json:"id,omitempty"`
}

func (e Event) String() string {
	return Stringify(e)
}

func toTimestamp(t interface{}) string {
	switch t.(type) {
		case json.Number: {
			i, err := t.(json.Number).Int64()
			if err != nil {
				panic(fmt.Sprintf("error deserializing number: %v, error: %v", t, err))
			}
			tm := time.Unix(i, 0)
			return tm.Format("2006-01-02T15:04:05Z")
		}
		case float64: {
			tm := time.Unix(int64(t.(float64)), 0)
			return tm.Format("2006-01-02T15:04:05Z")
		}
	}
	return t.(string)
}

func applyPushEventHack(buf []byte) []byte {
	// webhooks have an issue where the pushed_at and created_at come in as
	// integers vs string like the normal API so we need to fix them before we
	// deserialize them or we get an error
	var p map[string]interface{}
	d := json.NewDecoder(strings.NewReader(string(buf)))
	d.UseNumber() // prevent numbers from getting converted
	if err := d.Decode(&p); err != nil {
		panic(err.Error())
	}
	r := p["repository"].(map[string]interface{})
	r["pushed_at"] = toTimestamp(r["pushed_at"])
	r["created_at"] = toTimestamp(r["created_at"])
	result, err := json.Marshal(p)
	if err != nil {
		panic(err.Error())
	}
	return result
}

// Payload returns the parsed event payload. For recognized event types,
// a value of the corresponding struct type will be returned.
func (e *Event) Payload() (payload interface{}) {
	switch *e.Type {
	case "CommitCommentEvent":
		payload = &CommitCommentEvent{}
	case "CreateEvent":
		payload = &CreateEvent{}
	case "DeleteEvent":
		payload = &DeleteEvent{}
	case "DeploymentEvent":
		payload = &DeploymentEvent{}
	case "DeploymentStatusEvent":
		payload = &DeploymentStatusEvent{}
	case "ForkEvent":
		payload = &ForkEvent{}
	case "GollumEvent":
		payload = &GollumEvent{}
	// case "IntegrationInstallationEvent":
	// 	payload = &IntegrationInstallationEvent{}
	// case "IntegrationInstallationRepositoriesEvent":
	// 	payload = &IntegrationInstallationRepositoriesEvent{}
	// case "IssueActivityEvent":
	// 	payload = &IssueActivityEvent{}
	case "IssueCommentEvent":
		payload = &IssueCommentEvent{}
	case "IssuesEvent":
		payload = &IssuesEvent{}
	case "MemberEvent":
		payload = &MemberEvent{}
	case "MembershipEvent":
		payload = &MembershipEvent{}
	case "PageBuildEvent":
		payload = &PageBuildEvent{}
	case "PublicEvent":
		payload = &PublicEvent{}
	case "PullRequestEvent":
		payload = &PullRequestEvent{}
	case "PullRequestReviewCommentEvent":
		payload = &PullRequestReviewCommentEvent{}
	case "PushEvent":
		*e.RawPayload = applyPushEventHack(*e.RawPayload)
		payload = &PushEvent{}
	case "ReleaseEvent":
		payload = &ReleaseEvent{}
	case "RepositoryEvent":
		payload = &RepositoryEvent{}
	case "StatusEvent":
		payload = &StatusEvent{}
	case "TeamAddEvent":
		payload = &TeamAddEvent{}
	case "WatchEvent":
		payload = &WatchEvent{}
	}
	if err := json.Unmarshal(*e.RawPayload, &payload); err != nil {
		panic(err.Error())
	}
	return payload
}
