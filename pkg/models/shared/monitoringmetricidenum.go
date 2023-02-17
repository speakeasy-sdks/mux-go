package shared

type MonitoringMetricIDEnum string

const (
	MonitoringMetricIDEnumCurrentConcurrentViewers     MonitoringMetricIDEnum = "current-concurrent-viewers"
	MonitoringMetricIDEnumCurrentRebufferingPercentage MonitoringMetricIDEnum = "current-rebuffering-percentage"
	MonitoringMetricIDEnumExitsBeforeVideoStart        MonitoringMetricIDEnum = "exits-before-video-start"
	MonitoringMetricIDEnumPlaybackFailurePercentage    MonitoringMetricIDEnum = "playback-failure-percentage"
	MonitoringMetricIDEnumCurrentAverageBitrate        MonitoringMetricIDEnum = "current-average-bitrate"
)
