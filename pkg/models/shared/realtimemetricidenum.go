package shared

type RealtimeMetricIDEnum string

const (
	RealtimeMetricIDEnumCurrentConcurrentViewers     RealtimeMetricIDEnum = "current-concurrent-viewers"
	RealtimeMetricIDEnumCurrentRebufferingPercentage RealtimeMetricIDEnum = "current-rebuffering-percentage"
	RealtimeMetricIDEnumExitsBeforeVideoStart        RealtimeMetricIDEnum = "exits-before-video-start"
	RealtimeMetricIDEnumPlaybackFailurePercentage    RealtimeMetricIDEnum = "playback-failure-percentage"
	RealtimeMetricIDEnumCurrentAverageBitrate        RealtimeMetricIDEnum = "current-average-bitrate"
)
