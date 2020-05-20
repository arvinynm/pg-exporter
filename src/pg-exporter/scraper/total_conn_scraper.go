package scraper

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"pg-exporter/log"
)

const totalConnectionsSql = "select count(*) from pg_stat_activity"

type TotalConnectScraper struct {}

func NewTotalConnectScraper () *TotalConnectScraper {
	return &TotalConnectScraper{}
}

func (t *TotalConnectScraper) Scrape(db *sql.DB, ch chan<- prometheus.Metric) error {
	var (
		err error
		count float64
	)

	rows, err := db.Query(totalConnectionsSql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return err
		}
		log.Infof("total connections count: %f", count)
	}

	totalConnectGauge.Set(count)
	ch <- totalConnectGauge
	return nil
}

func (t *TotalConnectScraper) Name() string {
	return "total_connections_scraper"
}

var (
	totalConnOpt = prometheus.GaugeOpts{
		Name: "total_connections",
		Help: "postgres total connections",
		ConstLabels:prometheus.Labels{"zone": "postgres"},
	}
	totalConnectGauge = prometheus.NewGauge(totalConnOpt)
)
