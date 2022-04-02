package file_mon

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
    mb.Registry.MustAddMetricSet("my_module", "file_mon", New)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
	FileConfig 				[]FileConfig
	file_name 				string
	max_delta 				int
	default_max_delta		int
	default_start_time		[]int
	default_end_time		[]int
	default_week_days		[]int
	delta 					int
	alert 					bool
	active					bool
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
		BaseMetricSet:  		base,
		FileConfig:     		config.FileConfig,
		default_max_delta:  	config.default_max_delta,
		default_start_time:		config.default_start_time,
		default_end_time:		config.default_end_time,
		default_week_days:		config.default_week_days,
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

	for _, file_config := range FileConfig{

		
		f, _ := os.Open(file_config.FileName)
		out, _ := f.Stat()
		mod_time := out.ModTime()
		delta := act_time.Sub(mod_time).Seconds()
		m.delta = int(math.Round(delta))
		m.file_name = file_config.FileName
		m.alert = false
		active := false
		week_days := file_config.week_days
		start_time := file_config.start_time
		end_time := file_config.end_time
		// read todays weekday
		act_weekday := int(time.Now().Weekday())
		// read if monitoring_week_days is set, if not set default value
		if week_days == nil {
			week_days = m.default_week_days
		}
		for _, x := range week_days {
			if act_weekday == x {
				active = true
			}
		}
		// read if max_delta is set, if not, add default value
		m.max_delta = file_config.MaxDelta
		// read if monitorin_start_time is set, if not set default value
		if start_time == nil {
			start_time = m.default_start_time
		}
		// read if monitoring_end_time is set, if not set default value
		if end_time == nil {
			end_time = m.default_end_time
		}
		// evaluate if we are now in a monitoring time window
		window_start := time.Date(year, month, start_time[0], start_time[1], 0, 0, 0, time.UTC)
		window_end := time.Date(year, month, end_time[0], end_time[1], 0, 0, 0, time.UTC)
		if window_start.After(act_time) && window_end.Before(act_time) && active {
			// evaluate if this is an alert situation
			if file_config.MaxDelta < m.delta {
				m.alert = true
			}
		} 

		report.Event(mb.Event{
			MetricSetFields: common.MapStr{
				"delta": m.delta,
				"max_delta": file_config.MaxDelta,
				"file_name": file_config.FileName,
				"alert": m.alert,
			},
		})


	}

	return nil
}

// type DefaultConfig struct {
// 	default_max_delta				int		`config:"default_max_delta"`
// 	default_start_time				string	`config:"default_start_time"`
// 	default_end_time				string	`config:"default_end_time"`
// 	default_monitoring_week_days 	[]int	`config:"default_monitoring_week_days"`
// }

type FileConfig struct {
	FileName			string	`config:"file_name"`
	MaxDelta			int		`config:"max_delta"`
	start_time			[]int	`config:"start_time"`
	end_time			[]int	`config:"end_time"`
	week_days			[]int	`config:"week_days"`
}

type Config struct {
	FileConfig      	[]FileConfig 	`config:"files"`
	default_max_delta	int		`config:"default_max_delta"`
	default_start_time	[]int	`config:"default_start_time"`
	default_end_time	[]int	`config:"default_end_time"`
	default_week_days 	[]int	`config:"default_week_days"`
}

func returnConfig() Config {
	return Config{}
}

