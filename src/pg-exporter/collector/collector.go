package collector

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"pg-exporter/log"
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
	for _, s := range p.Scrapers {
		err := s.Scrape(p.DB, ch)
		if err != nil {
			log.Error(err.Error())
		}
	}
}