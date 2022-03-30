package cdaxbeat 

import (
        "sync"

        "github.com/elastic/beats/v7/metricbeat/internal/sysinit"
        "github.com/elastic/beats/v7/metricbeat/mb"
)

var once sync.Once

func init() {
        // Register the ModuleFactory function for the "system" module.
        if err := mb.Registry.AddModule("cdaxbeat", sysinit.InitSystemModule); err != nil {
                panic(err)
        }
}

