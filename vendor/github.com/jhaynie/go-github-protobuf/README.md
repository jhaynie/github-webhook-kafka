# Go Bindings for Github API protobuf

This project contains the Go bindings for the [protobuf](https://github.com/google/protobuf) definition files for entities and events in the [Github API](https://developer.github.com/).  This bindings generated from the [jhaynie/github-protobuf](https://github.com/jhaynie/github-protobuf) project.

## Install

```go
import "github.com/jhaynie/go-github-protobuf/github"
```

## Usage

Parse an incoming Github WebHook event and convert to the appropriate event class:

```go
// r is the http request
payload, err := github.ValidatePayload(r, "your secret")
if err != nil {
	// error
}
event, err := github.ParseWebHook(github.WebHookType(r), payload)
if err != nil {
	// handle error
}
switch event := event.(type) {
	case CommitCommentEvent:
	    processCommitCommentEvent(event)
	case CreateEvent:
	    processCreateEvent(event)
}
```


## License

Licensed under the MIT License.
