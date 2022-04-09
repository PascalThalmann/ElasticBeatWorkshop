package file_mon

// https://github.com/elastic/beats/blob/master/docs/devguide/create-metricset.asciidoc#creating-metricbeat-module

import (
	"math"
	"os"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("my_module", "file_mon", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	DefaultMaxDelta  int          `config:"default_max_delta"`
	DefaultStartTime []int        `config:"default_start_time"`
	DefaultEndTime   []int        `config:"default_end_time"`
	DefaultWeekDays  []int        `config:"default_week_days"`
	FileConfig       []FileConfig `config:"files"`
}

type FileConfig struct {
	FileName  string `config:"file_name"`
	MaxDelta  int    `config:"max_delta"`
	StartTime []int  `config:"start_time"`
	EndTime   []int  `config:"end_time"`
	WeekDays  []int  `config:"week_days"`
}

func returnConfig() MetricSet {
	return MetricSet{}
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The my_module file_mon metricset is beta.")

	config := returnConfig()

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	return &MetricSet{
		BaseMetricSet:    base,
		FileConfig:       config.FileConfig,
		DefaultMaxDelta:  config.DefaultMaxDelta,
		DefaultStartTime: config.DefaultStartTime,
		DefaultEndTime:   config.DefaultEndTime,
		DefaultWeekDays:  config.DefaultWeekDays,
	}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {

	FileConfig := m.FileConfig

	act_time := time.Now()
	year := act_time.Year()
	month := act_time.Month()

	for _, file_config := range FileConfig {

		f, err := os.Open(file_config.FileName)

		// we check if the file is even exisisting
		if err != nil {
			report.Event(mb.Event{
				MetricSetFields: common.MapStr{
					"error":     "file not existing",
					"file_name": file_config.FileName,
					"warning":   true,
				},
			})
			continue
		}
		out, _ := f.Stat()
		mod_time := out.ModTime()
		difference := act_time.Sub(mod_time).Seconds()
		delta := int(math.Round(difference))
		alert := false
		active := false

		// read todays weekday
		act_weekday := int(time.Now().Weekday())

		// read if monitoring_week_days is set, if not set default value
		if len(file_config.WeekDays) == 0 {
			file_config.WeekDays = m.DefaultWeekDays
		}
		for _, x := range file_config.WeekDays {
			if act_weekday == x {
				active = true
			}
		}

		// read if max_delta is set. if not, add default value
		if file_config.MaxDelta == 0 {
			file_config.MaxDelta = m.DefaultMaxDelta
		}

		// read if monitoring_start_time is set. if not, set default value
		if len(file_config.StartTime) == 0 {
			file_config.StartTime = m.DefaultStartTime
		}

		// read if monitoring_end_time is set. if not, set default value
		if len(file_config.EndTime) == 0 {
			file_config.EndTime = m.DefaultEndTime
		}

		// evaluate if we are now in a monitoring time window
		window_start := time.Date(year, month, file_config.StartTime[0], file_config.StartTime[1], 0, 0, 0, time.UTC)
		window_end := time.Date(year, month, file_config.EndTime[0], file_config.EndTime[1], 0, 0, 0, time.UTC)
		if window_start.Before(act_time) && window_end.After(act_time) && active {
			// evaluate alert
			if file_config.MaxDelta < delta {
				alert = true
			}
		} else {
			active = false
		}

		report.Event(mb.Event{
			MetricSetFields: common.MapStr{
				"delta":      delta,
				"max_delta":  file_config.MaxDelta,
				"file_name":  file_config.FileName,
				"alert":      alert,
				"active":     active,
				"start_time": file_config.StartTime,
				"end_time":   file_config.EndTime,
				"week_days":  file_config.WeekDays,

				"default_max_delta":  m.DefaultMaxDelta,
				"default_start_time": m.DefaultStartTime,
				"default_end_time":   m.DefaultEndTime,
				"default_week_days":  m.DefaultWeekDays,
			},
		})

	}

	return nil
}
