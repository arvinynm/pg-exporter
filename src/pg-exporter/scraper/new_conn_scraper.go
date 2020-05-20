package scraper

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

const newConnectionSql = "select count(*) from pg_stat_activity where now()-backend_start > '5 second'"

type NewConnectScraper struct {
	Query string
}

func NewNewConnectScraper () *NewConnectScraper {
	return &NewConnectScraper{}
}

func (t *NewConnectScraper) Scrape(db *sql.DB, ch chan<- prometheus.Metric) {
	var (
		err error
		count float64
	)

	rows, err := db.Query(newConnectionSql)
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

	newConnect30sGauge.Set(count)
	ch <- newConnect30sGauge
}

func (t *NewConnectScraper) Name() string {
	return "new_connections_in_30s_scraper"
}

var (
	newConnIn30sOpt = prometheus.GaugeOpts{
		Name: "new_connections_in_30s",
		Help: "postgres new connections in 30s",
		ConstLabels:prometheus.Labels{"zone": "postgres"},
	}
	newConnect30sGauge = prometheus.NewGauge(newConnIn30sOpt)
)