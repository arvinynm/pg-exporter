package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"pg-exporter/collector"
	"pg-exporter/scraper"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "current temperature of the cpu",
	})
	hdFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors",
	},
		[]string{"device"},
	)
)

func init() {
	var p = &collector.PgCollector{}
	p.Scrapers = append(p.Scrapers, &scraper.YnmPowerScraper{})

	prometheus.MustRegister(p)
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main() {
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device":"/dev/sda"}).Inc()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9096", nil))
}
