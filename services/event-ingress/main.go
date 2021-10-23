package main

import (
	"context"
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
)

func main() {
	ctx := context.Background()
	p, _ := cloudevents.NewHTTP()

	c, _ := cloudevents.NewClientHTTP(cehttp.WithTarget("http://broker-ingress.knative-eventing.svc.cluster.local/default/default"))

	h, _ := cloudevents.NewHTTPReceiveHandler(ctx, p, c.Send)

	log.Printf("will listen on :8080\n")
	if err := http.ListenAndServe(":8080", h); err != nil {
		log.Fatalf("unable to start http server, %s", err)
	}
}
