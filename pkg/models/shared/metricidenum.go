package shared

type MetricIDEnum string

const (
	MetricIDEnumAggregateStartupTime           MetricIDEnum = "aggregate_startup_time"
	MetricIDEnumDownscalePercentage            MetricIDEnum = "downscale_percentage"
	MetricIDEnumExitsBeforeVideoStart          MetricIDEnum = "exits_before_video_start"
	MetricIDEnumMaxDownscalePercentage         MetricIDEnum = "max_downscale_percentage"
	MetricIDEnumMaxUpscalePercentage           MetricIDEnum = "max_upscale_percentage"
	MetricIDEnumPageLoadTime                   MetricIDEnum = "page_load_time"
	MetricIDEnumPlaybackFailurePercentage      MetricIDEnum = "playback_failure_percentage"
	MetricIDEnumPlaybackFailureScore           MetricIDEnum = "playback_failure_score"
	MetricIDEnumPlayerStartupTime              MetricIDEnum = "player_startup_time"
	MetricIDEnumPlayingTime                    MetricIDEnum = "playing_time"
	MetricIDEnumRebufferCount                  MetricIDEnum = "rebuffer_count"
	MetricIDEnumRebufferDuration               MetricIDEnum = "rebuffer_duration"
	MetricIDEnumRebufferFrequency              MetricIDEnum = "rebuffer_frequency"
	MetricIDEnumRebufferPercentage             MetricIDEnum = "rebuffer_percentage"
	MetricIDEnumRebufferScore                  MetricIDEnum = "rebuffer_score"
	MetricIDEnumRequestsForFirstPreroll        MetricIDEnum = "requests_for_first_preroll"
	MetricIDEnumSeekLatency                    MetricIDEnum = "seek_latency"
	MetricIDEnumStartupTimeScore               MetricIDEnum = "startup_time_score"
	MetricIDEnumUniqueViewers                  MetricIDEnum = "unique_viewers"
	MetricIDEnumUpscalePercentage              MetricIDEnum = "upscale_percentage"
	MetricIDEnumVideoQualityScore              MetricIDEnum = "video_quality_score"
	MetricIDEnumVideoStartupPrerollLoadTime    MetricIDEnum = "video_startup_preroll_load_time"
	MetricIDEnumVideoStartupPrerollRequestTime MetricIDEnum = "video_startup_preroll_request_time"
	MetricIDEnumVideoStartupTime               MetricIDEnum = "video_startup_time"
	MetricIDEnumViewerExperienceScore          MetricIDEnum = "viewer_experience_score"
	MetricIDEnumViews                          MetricIDEnum = "views"
	MetricIDEnumWeightedAverageBitrate         MetricIDEnum = "weighted_average_bitrate"
)
