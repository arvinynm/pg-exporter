package scraper

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"pg-exporter/log"
)

const newConnectionSql = "select count(*) from pg_stat_activity where now()-backend_start > '5 second'"

type NewConnectScraper struct {
	Query string
}

func NewNewConnectScraper () *NewConnectScraper {
	return &NewConnectScraper{}
}

func (t *NewConnectScraper) Scrape(db *sql.DB, ch chan<- prometheus.Metric) error{
	var (
		err error
		count float64
	)

	rows, err := db.Query(newConnectionSql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return err
		}
		log.Infof("new connection count in 30s: %f", count)
	}

	newConnect30sGauge.Set(count)
	ch <- newConnect30sGauge
	return nil
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