package scraper

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type TotalConnectionScraper struct {
}

func (t *TotalConnectionScraper) Scrape(ch chan<- prometheus.Metric) {

}

type TotalConnectScraper struct {
	Query string
}

func (t *TotalConnectScraper) Scrape(db *sql.DB, ch chan<- prometheus.Metric) {
	var (
		err error
		count float64
	)

	rows, err := db.Query(t.Query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(count)
	}

	totalConnectGauge.Set(count)
	ch <- totalConnectGauge
}

func (t *TotalConnectScraper) Name() string {
	return "ynm_power_scraper"
}

var (
	gaugeOpt = prometheus.GaugeOpts{
		Name: "total_connections",
		Help: "postgres total connections",
		ConstLabels:prometheus.Labels{"zone": "postgres"},
	}
	totalConnectGauge = prometheus.NewGauge(gaugeOpt)
)