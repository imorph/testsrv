package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
)

func main() {
	http.HandleFunc("/", handle(home))
	http.HandleFunc("/articles", handle(articles))
	http.HandleFunc("/metrics", handle(metricsPage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func articles(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Get("https://www.google.com/search?q=ansible+articles")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}

func handle(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := fmt.Sprintf(`requests_total{path=%q}`, r.URL.Path)
		metrics.GetOrCreateCounter(s).Inc()
		h.ServeHTTP(w, r)
	}
}

func metricsPage(w http.ResponseWriter, r *http.Request) {
	metrics.WritePrometheus(w, true)
}
