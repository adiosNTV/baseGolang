package cache

import (
	"baseGo/internal/pkg/jsonlog"
	"github.com/maypok86/otter"
	"os"
)

var Store otter.CacheWithVariableTTL[string, string]

func init() {
	var err error
	Store, err = otter.MustBuilder[string, string](10_000).
		CollectStats().
		Cost(func(key string, value string) uint32 {
			return 1
		}).
		WithVariableTTL().
		Build()
	if err != nil {
		logr := jsonlog.NewLogger(os.Stdout, jsonlog.LevelInfo)
		logr.PrintFatal(err, nil)
	}
}
