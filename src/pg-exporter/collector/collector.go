package collector

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"pg-exporter/scraper"
)

type PgCollector struct {
	Scrapers []scraper.Scraper
	DB *sql.DB
}


func (p *PgCollector) Describe(ch chan<- *prometheus.Desc) {
	//ch <- gauge.Desc()
}

func (p *PgCollector) Collect(ch chan<- prometheus.Metric) {
	p.Scrape(ch)
}


func (p *PgCollector) Scrape (ch chan <-prometheus.Metric) {
	fmt.Println("coming here")
	for _, s := range p.Scrapers {
		s.Scrape(p.DB, ch)
	}
}