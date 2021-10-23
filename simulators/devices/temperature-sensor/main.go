package main

import (
	"context"
	"fmt"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cehttp "github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type TemperatureLowAlarm struct {
	Val float32
}

var temperature = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Name:        "temperature",
	Help:        "The duration of the last loop of the status update controller.",
	ConstLabels: prometheus.Labels{"unit": "c"},
}, []string{
	"name",
})

func init() {
	prometheus.MustRegister(temperature)
}

func main() {

	c, _ := cloudevents.NewClientHTTP(cehttp.WithTarget("http://event-ingress.default.127.0.0.1.nip.io/"))

	e := cloudevents.NewEvent()

	e.SetID(uuid.New().String())
	e.SetSource("kates/simulators/devices/temperature-sensor")
	e.SetType("kates.citadel.works/temperature-low-alarm")

	e.SetData(cloudevents.ApplicationJSON, TemperatureLowAlarm{Val: 912.0})

	temperature.WithLabelValues("simulated-temperature").Set(12)

	res := c.Send(context.TODO(), e)

	fmt.Println(res.Error())

	http.Handle("/metrics", promhttp.Handler())
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
