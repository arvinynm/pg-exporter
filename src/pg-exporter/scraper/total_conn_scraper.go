package scraper

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

const maxConnectionsSql = "select count(*) from pg_stat_activity"

type TotalConnectScraper struct {}

func NewTotalConnectScraper () *TotalConnectScraper {
	return &TotalConnectScraper{}
}

func (t *TotalConnectScraper) Scrape(db *sql.DB, ch chan<- prometheus.Metric) {
	var (
		err error
		count float64
	)

	rows, err := db.Query(maxConnectionsSql)
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
