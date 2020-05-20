package scraper

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type Scraper interface {
	Scrape(db *sql.DB, ch chan<- prometheus.Metric) error
	Name() string
}
