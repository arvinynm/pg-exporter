package scraper

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type TotalConnectionScraper struct {
}

func (t *TotalConnectionScraper) Scrape(ch chan<- prometheus.Metric) {

}

type YnmPowerScraper struct {

}

func (y *YnmPowerScraper) Scrape(ch chan<- prometheus.Metric) {
	ynmPowerGauge.Set(index)
	ch <- ynmPowerGauge
}

func (y *YnmPowerScraper) Name() string {
	return "ynm_power_scraper"
}

var (
	index float64
	gaugeOpt = prometheus.GaugeOpts{
		Name: "ynm_dex",
		Help: "yanningmin's dex",
		ConstLabels:prometheus.Labels{"zone": "ynm"},
	}
	ynmPowerGauge = prometheus.NewGauge(gaugeOpt)
)


func init() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			index++
		}
	}()
}