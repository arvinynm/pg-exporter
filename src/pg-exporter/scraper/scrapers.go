package scraper

import "github.com/prometheus/client_golang/prometheus"

type Scraper interface {
	Scrape(ch chan<- prometheus.Metric)
	Name() string
}
