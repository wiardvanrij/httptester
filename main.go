package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tcnksm/go-httpstat"
)

var (
	requestDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "http_stats",
			Help:       "Http statistics",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"url", "type", "status"},
	)
)

func main() {

	seconds := flag.Int("seconds", 1, "seconds between each request")
	flag.Parse()

	prometheus.MustRegister(requestDuration)
	urls := flag.Args()

	done := make(chan bool)
	ts := make([]*time.Ticker, len(urls))
	for i, url := range urls {
		ts[i] = schedule(url, time.Duration(*seconds)*time.Second, done)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

func schedule(url string, interval time.Duration, done <-chan bool) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				doRequest(url)
			case <-done:
				return
			}
		}
	}()
	return ticker
}

func doRequest(url string) {

	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create go-httpstat powered context and pass it to http.Request
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)

	} else {
		if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
			log.Println(err)
		}

		res.Body.Close()
		result.End(time.Now())

		requestDuration.WithLabelValues(url, "Connect", strconv.Itoa(res.StatusCode)).Observe(float64(result.Connect / time.Millisecond))
		requestDuration.WithLabelValues(url, "DNSLookup", strconv.Itoa(res.StatusCode)).Observe(float64(result.DNSLookup / time.Millisecond))
		requestDuration.WithLabelValues(url, "NameLookup", strconv.Itoa(res.StatusCode)).Observe(float64(result.NameLookup / time.Millisecond))
		requestDuration.WithLabelValues(url, "Pretransfer", strconv.Itoa(res.StatusCode)).Observe(float64(result.Pretransfer / time.Millisecond))
		requestDuration.WithLabelValues(url, "ServerProcessing", strconv.Itoa(res.StatusCode)).Observe(float64(result.ServerProcessing / time.Millisecond))
		requestDuration.WithLabelValues(url, "StartTransfer", strconv.Itoa(res.StatusCode)).Observe(float64(result.StartTransfer / time.Millisecond))
		requestDuration.WithLabelValues(url, "TCPConnection", strconv.Itoa(res.StatusCode)).Observe(float64(result.TCPConnection / time.Millisecond))
		requestDuration.WithLabelValues(url, "TLSHandshake", strconv.Itoa(res.StatusCode)).Observe(float64(result.TLSHandshake / time.Millisecond))
		requestDuration.WithLabelValues(url, "Total", strconv.Itoa(res.StatusCode)).Observe(float64(time.Since(start) / time.Millisecond))
	}

}
