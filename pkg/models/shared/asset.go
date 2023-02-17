package shared

import (
	"time"
)

// AssetErrors
// Object that describes any errors that happened when processing this asset.
type AssetErrors struct {
	Messages []string `json:"messages,omitempty"`
	Type     *string  `json:"type,omitempty"`
}

type AssetMasterStatusEnum string

const (
	AssetMasterStatusEnumReady     AssetMasterStatusEnum = "ready"
	AssetMasterStatusEnumPreparing AssetMasterStatusEnum = "preparing"
	AssetMasterStatusEnumErrored   AssetMasterStatusEnum = "errored"
)

// AssetMaster
// An object containing the current status of Master Access and the link to the Master MP4 file when ready. This object does not exist if `master_acess` is set to `none` and when the temporary URL expires.
type AssetMaster struct {
	Status *AssetMasterStatusEnum `json:"status,omitempty"`
	URL    *string                `json:"url,omitempty"`
}

type AssetMasterAccessEnum string

const (
	AssetMasterAccessEnumTemporary AssetMasterAccessEnum = "temporary"
	AssetMasterAccessEnumNone      AssetMasterAccessEnum = "none"
)

type AssetMaxStoredResolutionEnum string

const (
	AssetMaxStoredResolutionEnumAudioOnly AssetMaxStoredResolutionEnum = "Audio only"
	AssetMaxStoredResolutionEnumSd        AssetMaxStoredResolutionEnum = "SD"
	AssetMaxStoredResolutionEnumHd        AssetMaxStoredResolutionEnum = "HD"
	AssetMaxStoredResolutionEnumFhd       AssetMaxStoredResolutionEnum = "FHD"
	AssetMaxStoredResolutionEnumUhd       AssetMaxStoredResolutionEnum = "UHD"
)

type AssetMp4SupportEnum string

const (
	AssetMp4SupportEnumStandard AssetMp4SupportEnum = "standard"
	AssetMp4SupportEnumNone     AssetMp4SupportEnum = "none"
)

type AssetNonStandardInputReasonsAudioEditListEnum string

const (
	AssetNonStandardInputReasonsAudioEditListEnumNonStandard AssetNonStandardInputReasonsAudioEditListEnum = "non-standard"
)

type AssetNonStandardInputReasonsUnexpectedMediaFileParametersEnum string

const (
	AssetNonStandardInputReasonsUnexpectedMediaFileParametersEnumNonStandard AssetNonStandardInputReasonsUnexpectedMediaFileParametersEnum = "non-standard"
)

type AssetNonStandardInputReasonsVideoBitrateEnum string

const (
	AssetNonStandardInputReasonsVideoBitrateEnumHigh AssetNonStandardInputReasonsVideoBitrateEnum = "high"
)

type AssetNonStandardInputReasonsVideoEditListEnum string

const (
	AssetNonStandardInputReasonsVideoEditListEnumNonStandard AssetNonStandardInputReasonsVideoEditListEnum = "non-standard"
)

type AssetNonStandardInputReasonsVideoGopSizeEnum string

const (
	AssetNonStandardInputReasonsVideoGopSizeEnumHigh AssetNonStandardInputReasonsVideoGopSizeEnum = "high"
)

// AssetNonStandardInputReasons
// An object containing one or more reasons the input file is non-standard. See [the guide on minimizing processing time](https://docs.mux.com/guides/video/minimize-processing-time) for more information on what a standard input is defined as. This object only exists on on-demand assets that have non-standard inputs, so if missing you can assume the input qualifies as standard.
type AssetNonStandardInputReasons struct {
	AudioCodec                    *string                                                        `json:"audio_codec,omitempty"`
	AudioEditList                 *AssetNonStandardInputReasonsAudioEditListEnum                 `json:"audio_edit_list,omitempty"`
	PixelAspectRatio              *string                                                        `json:"pixel_aspect_ratio,omitempty"`
	UnexpectedMediaFileParameters *AssetNonStandardInputReasonsUnexpectedMediaFileParametersEnum `json:"unexpected_media_file_parameters,omitempty"`
	VideoBitrate                  *AssetNonStandardInputReasonsVideoBitrateEnum                  `json:"video_bitrate,omitempty"`
	VideoCodec                    *string                                                        `json:"video_codec,omitempty"`
	VideoEditList                 *AssetNonStandardInputReasonsVideoEditListEnum                 `json:"video_edit_list,omitempty"`
	VideoFrameRate                *string                                                        `json:"video_frame_rate,omitempty"`
	VideoGopSize                  *AssetNonStandardInputReasonsVideoGopSizeEnum                  `json:"video_gop_size,omitempty"`
	VideoResolution               *string                                                        `json:"video_resolution,omitempty"`
}

type AssetRecordingTimesTypeEnum string

const (
	AssetRecordingTimesTypeEnumContent AssetRecordingTimesTypeEnum = "content"
	AssetRecordingTimesTypeEnumSlate   AssetRecordingTimesTypeEnum = "slate"
)

type AssetRecordingTimes struct {
	Duration  *float64                     `json:"duration,omitempty"`
	StartedAt *time.Time                   `json:"started_at,omitempty"`
	Type      *AssetRecordingTimesTypeEnum `json:"type,omitempty"`
}

type AssetStaticRenditionsFilesExtEnum string

const (
	AssetStaticRenditionsFilesExtEnumMp4 AssetStaticRenditionsFilesExtEnum = "mp4"
	AssetStaticRenditionsFilesExtEnumM4a AssetStaticRenditionsFilesExtEnum = "m4a"
)

type AssetStaticRenditionsFilesNameEnum string

const (
	AssetStaticRenditionsFilesNameEnumLowMp4    AssetStaticRenditionsFilesNameEnum = "low.mp4"
	AssetStaticRenditionsFilesNameEnumMediumMp4 AssetStaticRenditionsFilesNameEnum = "medium.mp4"
	AssetStaticRenditionsFilesNameEnumHighMp4   AssetStaticRenditionsFilesNameEnum = "high.mp4"
	AssetStaticRenditionsFilesNameEnumAudioM4a  AssetStaticRenditionsFilesNameEnum = "audio.m4a"
)

type AssetStaticRenditionsFiles struct {
	Bitrate  *int64                              `json:"bitrate,omitempty"`
	Ext      *AssetStaticRenditionsFilesExtEnum  `json:"ext,omitempty"`
	Filesize *string                             `json:"filesize,omitempty"`
	Height   *int                                `json:"height,omitempty"`
	Name     *AssetStaticRenditionsFilesNameEnum `json:"name,omitempty"`
	Width    *int                                `json:"width,omitempty"`
}

type AssetStaticRenditionsStatusEnum string

const (
	AssetStaticRenditionsStatusEnumReady     AssetStaticRenditionsStatusEnum = "ready"
	AssetStaticRenditionsStatusEnumPreparing AssetStaticRenditionsStatusEnum = "preparing"
	AssetStaticRenditionsStatusEnumDisabled  AssetStaticRenditionsStatusEnum = "disabled"
	AssetStaticRenditionsStatusEnumErrored   AssetStaticRenditionsStatusEnum = "errored"
)

// AssetStaticRenditions
// An object containing the current status of any static renditions (mp4s). The object does not exist if no static renditions have been requested. See [Download your videos](https://docs.mux.com/guides/video/download-your-videos) for more information.
type AssetStaticRenditions struct {
	Files  []AssetStaticRenditionsFiles     `json:"files,omitempty"`
	Status *AssetStaticRenditionsStatusEnum `json:"status,omitempty"`
}

type AssetStatusEnum string

const (
	AssetStatusEnumPreparing AssetStatusEnum = "preparing"
	AssetStatusEnumReady     AssetStatusEnum = "ready"
	AssetStatusEnumErrored   AssetStatusEnum = "errored"
)

type Asset struct {
	AspectRatio             *string                       `json:"aspect_ratio,omitempty"`
	CreatedAt               *string                       `json:"created_at,omitempty"`
	Duration                *float64                      `json:"duration,omitempty"`
	Errors                  *AssetErrors                  `json:"errors,omitempty"`
	ID                      *string                       `json:"id,omitempty"`
	IsLive                  *bool                         `json:"is_live,omitempty"`
	LiveStreamID            *string                       `json:"live_stream_id,omitempty"`
	Master                  *AssetMaster                  `json:"master,omitempty"`
	MasterAccess            *AssetMasterAccessEnum        `json:"master_access,omitempty"`
	MaxStoredFrameRate      *float64                      `json:"max_stored_frame_rate,omitempty"`
	MaxStoredResolution     *AssetMaxStoredResolutionEnum `json:"max_stored_resolution,omitempty"`
	Mp4Support              *AssetMp4SupportEnum          `json:"mp4_support,omitempty"`
	NonStandardInputReasons *AssetNonStandardInputReasons `json:"non_standard_input_reasons,omitempty"`
	NormalizeAudio          *bool                         `json:"normalize_audio,omitempty"`
	Passthrough             *string                       `json:"passthrough,omitempty"`
	PerTitleEncode          *bool                         `json:"per_title_encode,omitempty"`
	PlaybackIds             []PlaybackID                  `json:"playback_ids,omitempty"`
	RecordingTimes          []AssetRecordingTimes         `json:"recording_times,omitempty"`
	SourceAssetID           *string                       `json:"source_asset_id,omitempty"`
	StaticRenditions        *AssetStaticRenditions        `json:"static_renditions,omitempty"`
	Status                  *AssetStatusEnum              `json:"status,omitempty"`
	Test                    *bool                         `json:"test,omitempty"`
	Tracks                  []Track                       `json:"tracks,omitempty"`
	UploadID                *string                       `json:"upload_id,omitempty"`
}
