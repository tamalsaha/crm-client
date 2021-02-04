package main

import (
	"encoding/json"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/mailgun/mailgun-go/v4/events"
	"net/http"
	"os"
	"sigs.k8s.io/yaml"
)

func main() {
	// Create an instance of the Mailgun Client
	mg, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		fmt.Printf("mailgun error: %s\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var payload mailgun.WebhookPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			fmt.Printf("decode JSON error: %s", err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		verified, err := mg.VerifyWebhookSignature(payload.Signature)
		if err != nil {
			fmt.Printf("verify error: %s\n", err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		if !verified {
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Printf("failed verification %+v\n", payload.Signature)
			return
		}

		fmt.Printf("Verified Signature\n")

		// Parse the raw event to extract the

		e, err := mailgun.ParseEvent(payload.EventData)
		if err != nil {
			fmt.Printf("parse event error: %s\n", err)
			return
		}

		d2, err := yaml.Marshal(e)
		if err != nil {
			fmt.Printf("parse event error: %s\n", err)
			return
		}

		switch event := e.(type) {
		case *events.Accepted:
			fmt.Printf("Accepted: auth: %t\n", event.Flags.IsAuthenticated)
		case *events.Delivered:
			fmt.Printf("Delivered transport: %s\n", event.Envelope.Transport)
		case *events.Opened:
			fmt.Printf("Delivered transport: %s\n", string(d2))
		}
	})

	fmt.Println("Running...")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("serve error: %s\n", err)
		os.Exit(1)
	}
}