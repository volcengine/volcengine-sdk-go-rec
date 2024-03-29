package metrics

import "time"

const (
	// TODO 根据region动态选择defaultMetricsDomain
	defaultMetricsDomain = "api.byteair.volces.com"
	defaultMetricsPrefix = "volcengine.rec.sdk"
	defaultHttpSchema    = "https"

	counterUrlFormat = "%s://%s/api/counter"
	otherUrlFormat   = "%s://%s/api/put"

	defaultFlushInterval = 10 * time.Second
	reservoirSize        = 65536
	decayAlpha           = 0.02
	maxTryTimes          = 2
	defaultHttpTimeout   = 800 * time.Millisecond

	delimiter = "+"
)

type metricsType int

const (
	metricsTypeCounter metricsType = iota
	metricsTypeTimer
	metricsTypeStore
)
