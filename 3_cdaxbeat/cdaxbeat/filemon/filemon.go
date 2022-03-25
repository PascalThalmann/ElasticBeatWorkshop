package filemon

// https://github.com/elastic/beats/blob/master/docs/devguide/create-metricset.asciidoc#creating-metricbeat-module

import (
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"math"
	"os"
	"time"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("cdaxbeat", "filemon", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	delta 		int
	files 		[]FileConfig
	file_name 	string
	max_delta 	int
	alert 		bool
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The {module} {metricset} metricset is beta.")
	config := returnConfig()

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet:	base,
		files:			config.Files,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {

	files := m.files

	for _, file_name := range files{

		act_time := time.Now()
		f, _ := os.Open(file_name.FileName)
		out, _ := f.Stat()
		mod_time := out.ModTime()
		delta := act_time.Sub(mod_time).Seconds()
		m.delta = int(math.Round(delta))
		m.file_name = file_name.FileName
		m.max_delta = file_name.MaxDelta
		m.alert = false

		if file_name.MaxDelta < m.delta {
			m.alert = true
		}

		report.Event(mb.Event{
			MetricSetFields: common.MapStr{
				"delta": m.delta,
				"max_delta": file_name.MaxDelta,
				"file_name": file_name.FileName,
				"alert": m.alert,
			},
		})

    }

	return nil
}

type Config struct {
    Files      []FileConfig  `config:"files"`
}

type FileConfig struct {
    FileName   string        `config:"file_name"`
    MaxDelta   int           `config:"max_delta"`
}

func returnConfig() Config {
    return Config{}
}