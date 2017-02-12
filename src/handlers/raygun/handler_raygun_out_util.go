package raygun

import (
	"github.com/grokify/glip-go-webhook"
)

func ExampleMessageGlip() (glipwebhook.GlipWebhookMessage, error) {
	msg, err := ExampleMessageSource()
	if err != nil {
		return glipwebhook.GlipWebhookMessage{}, err
	}
	return Normalize(msg), nil
}

func ExampleMessageSource() (RaygunOutMessage, error) {
	return RaygunOutMessageFromBytes(ExampleMessageBytes())
}

func ExampleMessageBytes() []byte {
	return []byte(`{
  "event":"error_notification",
  "eventType":"NewErrorOccurred",
  "error": {
    "url":"http://app.raygun.io/error-url",
    "message":"",
    "firstOccurredOn":"1970-01-28T01:49:36Z",
    "lastOccurredOn":"1970-01-28T01:49:36Z",
    "usersAffected":1,
    "totalOccurrences":1
  },
  "application": {
    "name":"application name",
    "url":"http://app.raygun.io/application-url"
  }
}`)
}