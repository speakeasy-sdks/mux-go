package shared

type GroupByEnum string

const (
	GroupByEnumAsn                      GroupByEnum = "asn"
	GroupByEnumBrowser                  GroupByEnum = "browser"
	GroupByEnumBrowserVersion           GroupByEnum = "browser_version"
	GroupByEnumCdn                      GroupByEnum = "cdn"
	GroupByEnumContinentCode            GroupByEnum = "continent_code"
	GroupByEnumCountry                  GroupByEnum = "country"
	GroupByEnumExitBeforeVideoStart     GroupByEnum = "exit_before_video_start"
	GroupByEnumExperimentName           GroupByEnum = "experiment_name"
	GroupByEnumOperatingSystem          GroupByEnum = "operating_system"
	GroupByEnumOperatingSystemVersion   GroupByEnum = "operating_system_version"
	GroupByEnumPlayerAutoplay           GroupByEnum = "player_autoplay"
	GroupByEnumPlayerErrorCode          GroupByEnum = "player_error_code"
	GroupByEnumPlayerMuxPluginName      GroupByEnum = "player_mux_plugin_name"
	GroupByEnumPlayerMuxPluginVersion   GroupByEnum = "player_mux_plugin_version"
	GroupByEnumPlayerName               GroupByEnum = "player_name"
	GroupByEnumPlayerPreload            GroupByEnum = "player_preload"
	GroupByEnumPlayerRemotePlayed       GroupByEnum = "player_remote_played"
	GroupByEnumPlayerSoftware           GroupByEnum = "player_software"
	GroupByEnumPlayerSoftwareVersion    GroupByEnum = "player_software_version"
	GroupByEnumPlayerVersion            GroupByEnum = "player_version"
	GroupByEnumPrerollAdAssetHostname   GroupByEnum = "preroll_ad_asset_hostname"
	GroupByEnumPrerollAdTagHostname     GroupByEnum = "preroll_ad_tag_hostname"
	GroupByEnumPrerollPlayed            GroupByEnum = "preroll_played"
	GroupByEnumPrerollRequested         GroupByEnum = "preroll_requested"
	GroupByEnumRegion                   GroupByEnum = "region"
	GroupByEnumSourceHostname           GroupByEnum = "source_hostname"
	GroupByEnumSourceType               GroupByEnum = "source_type"
	GroupByEnumStreamType               GroupByEnum = "stream_type"
	GroupByEnumSubPropertyID            GroupByEnum = "sub_property_id"
	GroupByEnumVideoEncodingVariant     GroupByEnum = "video_encoding_variant"
	GroupByEnumVideoID                  GroupByEnum = "video_id"
	GroupByEnumVideoSeries              GroupByEnum = "video_series"
	GroupByEnumVideoTitle               GroupByEnum = "video_title"
	GroupByEnumViewSessionID            GroupByEnum = "view_session_id"
	GroupByEnumViewerConnectionType     GroupByEnum = "viewer_connection_type"
	GroupByEnumViewerDeviceCategory     GroupByEnum = "viewer_device_category"
	GroupByEnumViewerDeviceManufacturer GroupByEnum = "viewer_device_manufacturer"
	GroupByEnumViewerDeviceModel        GroupByEnum = "viewer_device_model"
	GroupByEnumViewerDeviceName         GroupByEnum = "viewer_device_name"
	GroupByEnumViewerUserID             GroupByEnum = "viewer_user_id"
)
