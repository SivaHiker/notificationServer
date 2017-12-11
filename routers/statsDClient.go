package routers

import (
	"fmt"
	"log"

	"github.com/cactus/go-statsd-client/statsd"
)

type StatsDConfig struct {
	statterS statsd.Statter
}

var statsDConfigurator *StatsDConfig

func GetStatsDClient() *StatsDConfig {
	var (
		statsdClient1 statsd.Statter
		err           error
	)
	//var st StatsDConfig
	if statsDConfigurator == nil {
		// statsdClient1, err = statsd.NewClient("192.168.0.134:8125", "golang-client")
		statsdClient1, err = statsd.NewClient("10.15.0.38:7125", "golang-client")
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Successfully connected to StatsD Server ---->", statsdClient1)
			//st = StatsDConfig{statterS: statsdClient1}
			statsDConfigurator = &StatsDConfig{statterS: statsdClient1}
		}
	}
	return statsDConfigurator
}

func (StatsDConfig *StatsDConfig) CloseStatsDClient() {
	StatsDConfig.statterS.Close()
}

func (StatsDConfig *StatsDConfig) IncStat(stat string, value int64) {
	err := StatsDConfig.statterS.Inc(stat, value, 1.0)
	if err != nil {
		log.Fatal("there was an error sending the statsd inc stats", err)
	}
}

func (StatsDConfig *StatsDConfig) TimingStat(stat string, value int64) {
	err := StatsDConfig.statterS.Timing(stat, value, 1.0)
	if err != nil {
		log.Fatal("there was an error sending the statsd timing", err)
	}
}
