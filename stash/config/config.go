package config

import (
	"time"

	"github.com/tal-tech/go-queue/kq"
)

type (
	Condition struct {
		Key   string
		Value string
		Type  string `json:",default=match,options=match|contains"`
		Op    string `json:",default=and,options=and|or"`
	}

	ElasticSearchConf struct {
		Hosts            []string
		DailyIndexPrefix string
		TimeZone         string `json:",optional"`
		MaxChunkBytes    int    `json:",default=1048576"`
		Compress         bool   `json:",default=false"`
	}

	Filter struct {
		Action     string      `json:",options=drop|remove_field"`
		Conditions []Condition `json:",optional"`
		Fields     []string    `json:",optional"`
	}

	Processor struct {
		Input struct {
			Kafka kq.KqConf
		}
		Filters []Filter
		Output  struct {
			ElasticSearch ElasticSearchConf
		}
	}

	Config struct {
		Processors  []Processor
		GracePeriod time.Duration `json:",default=10s"`
	}
)
