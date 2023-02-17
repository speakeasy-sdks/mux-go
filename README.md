# github.com/speakeasy-sdks/mux-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/mux-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/mux-go"
    "github.com/speakeasy-sdks/mux-go/pkg/models/shared"
    "github.com/speakeasy-sdks/mux-go/pkg/models/operations"
)

func main() {
    opts := []mux.SDKOption{
        mux.WithSecurity(
            shared.Security{
                AccessToken: shared.SchemeAccessToken{
                    Password: "YOUR_PASSWORD_HERE",
                    Username: "YOUR_USERNAME_HERE",
                },
            }
        ),
    }

    s := mux.New(opts...)
    
    req := operations.CreateAssetRequest{
        Request: shared.CreateAssetRequest{
            Input: []shared.InputSettings{
                shared.InputSettings{
                    ClosedCaptions: false,
                    EndTime: 7151.9,
                    LanguageCode: "nulla",
                    Name: "id",
                    OverlaySettings: &shared.InputSettingsOverlaySettings{
                        Height: "vero",
                        HorizontalAlign: "right",
                        HorizontalMargin: "nulla",
                        Opacity: "nihil",
                        VerticalAlign: "bottom",
                        VerticalMargin: "facilis",
                        Width: "eum",
                    },
                    Passthrough: "iusto",
                    StartTime: 2975.34,
                    TextType: "undefined",
                    Type: "video",
                    URL: "sapiente",
                },
                shared.InputSettings{
                    ClosedCaptions: true,
                    EndTime: 3834.41,
                    LanguageCode: "voluptatum",
                    Name: "autem",
                    OverlaySettings: &shared.InputSettingsOverlaySettings{
                        Height: "vel",
                        HorizontalAlign: "right",
                        HorizontalMargin: "deleniti",
                        Opacity: "similique",
                        VerticalAlign: "middle",
                        VerticalMargin: "molestiae",
                        Width: "quo",
                    },
                    Passthrough: "quasi",
                    StartTime: 3373.96,
                    TextType: "subtitles",
                    Type: "text",
                    URL: "voluptatem",
                },
                shared.InputSettings{
                    ClosedCaptions: false,
                    EndTime: 8326.2,
                    LanguageCode: "a",
                    Name: "omnis",
                    OverlaySettings: &shared.InputSettingsOverlaySettings{
                        Height: "eos",
                        HorizontalAlign: "undefined",
                        HorizontalMargin: "accusamus",
                        Opacity: "reiciendis",
                        VerticalAlign: "middle",
                        VerticalMargin: "quibusdam",
                        Width: "et",
                    },
                    Passthrough: "praesentium",
                    StartTime: 5204.78,
                    TextType: "undefined",
                    Type: "text",
                    URL: "sed",
                },
            },
            MasterAccess: "undefined",
            Mp4Support: "standard",
            NormalizeAudio: true,
            Passthrough: "qui",
            PerTitleEncode: false,
            PlaybackPolicy: []shared.PlaybackPolicyEnum{
                "undefined",
                "signed",
                "public",
                "signed",
            },
            Test: true,
        },
    }
    
    res, err := s.Assets.CreateAsset(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AssetResponse != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations


### Assets

* `CreateAsset` - Create an asset
* `CreateAssetPlaybackID` - Create a playback ID
* `CreateAssetTrack` - Create an asset track
* `DeleteAsset` - Delete an asset
* `DeleteAssetPlaybackID` - Delete a playback ID
* `DeleteAssetTrack` - Delete an asset track
* `GetAsset` - Retrieve an asset
* `GetAssetInputInfo` - Retrieve asset input info
* `GetAssetPlaybackID` - Retrieve a playback ID
* `ListAssets` - List assets
* `UpdateAsset` - Update an Asset
* `UpdateAssetMasterAccess` - Update master access
* `UpdateAssetMp4Support` - Update MP4 support

### DeliveryUsage

* `ListDeliveryUsage` - List Usage

### Dimensions

* `ListDimensionValues` - Lists the values for a specific dimension
* `ListDimensions` - List Dimensions

### DirectUploads

* `CancelDirectUpload` - Cancel a direct upload
* `CreateDirectUpload` - Create a new direct upload URL
* `GetDirectUpload` - Retrieve a single direct upload's info
* `ListDirectUploads` - List direct uploads

### Errors

* `ListErrors` - List Errors

### Exports

* `ListExports` - List property video view export links
* `ListExportsViews` - List available property view exports

### Filters

* `ListFilterValues` - Lists values for a specific filter
* `ListFilters` - List Filters

### Incidents

* `GetIncident` - Get an Incident
* `ListIncidents` - List Incidents
* `ListRelatedIncidents` - List Related Incidents

### LiveStreams

* `CreateLiveStream` - Create a live stream
* `CreateLiveStreamPlaybackID` - Create a live stream playback ID
* `CreateLiveStreamSimulcastTarget` - Create a live stream simulcast target
* `DeleteLiveStream` - Delete a live stream
* `DeleteLiveStreamPlaybackID` - Delete a live stream playback ID
* `DeleteLiveStreamSimulcastTarget` - Delete a Live Stream Simulcast Target
* `DisableLiveStream` - Disable a live stream
* `EnableLiveStream` - Enable a live stream
* `GetLiveStream` - Retrieve a live stream
* `GetLiveStreamPlaybackID` - Retrieve a live stream playback ID
* `GetLiveStreamSimulcastTarget` - Retrieve a Live Stream Simulcast Target
* `ListLiveStreams` - List live streams
* `ResetStreamKey` - Reset a live stream's stream key
* `SignalLiveStreamComplete` - Signal a live stream is finished
* `UpdateLiveStream` - Update a live stream
* `UpdateLiveStreamEmbeddedSubtitles` - Update a live stream's embedded subtitles
* `UpdateLiveStreamGeneratedSubtitles` - Update a live stream's generated subtitles

### Metrics

* `GetMetricTimeseriesData` - Get metric timeseries data
* `GetOverallValues` - Get Overall values
* `ListAllMetricValues` - List all metric values
* `ListBreakdownValues` - List breakdown values
* `ListInsights` - List Insights

### Monitoring

* `GetMonitoringBreakdown` - Get Monitoring Breakdown
* `GetMonitoringHistogramTimeseries` - Get Monitoring Histogram Timeseries
* `GetMonitoringTimeseries` - Get Monitoring Timeseries
* `ListMonitoringDimensions` - List Monitoring Dimensions
* `ListMonitoringMetrics` - List Monitoring Metrics

### PlaybackID

* `GetAssetOrLivestreamID` - Retrieve an Asset or Live Stream ID

### PlaybackRestrictions

* `CreatePlaybackRestriction` - Create a Playback Restriction
* `DeletePlaybackRestriction` - Delete a Playback Restriction
* `GetPlaybackRestriction` - Retrieve a Playback Restriction
* `ListPlaybackRestrictions` - List Playback Restrictions
* `UpdateReferrerDomainRestriction` - Update the Referrer Playback Restriction

### RealTime

* `GetRealtimeBreakdown` - Get Real-Time Breakdown
* `GetRealtimeHistogramTimeseries` - Get Real-Time Histogram Timeseries
* `GetRealtimeTimeseries` - Get Real-Time Timeseries
* `ListRealtimeDimensions` - List Real-Time Dimensions
* `ListRealtimeMetrics` - List Real-Time Metrics

### Spaces

* `CreateSpace` - Create a space
* `CreateSpaceBroadcast` - Create a space broadcast
* `DeleteSpace` - Delete a space
* `DeleteSpaceBroadcast` - Delete a space broadcast
* `GetSpace` - Retrieve a space
* `GetSpaceBroadcast` - Retrieve space broadcast
* `ListSpaces` - List spaces
* `StartSpaceBroadcast` - Start a space broadcast
* `StopSpaceBroadcast` - Stop a space broadcast

### TranscriptionVocabularies

* `CreateTranscriptionVocabulary` - Create a Transcription Vocabulary
* `DeleteTranscriptionVocabulary` - Delete a Transcription Vocabulary
* `GetTranscriptionVocabulary` - Retrieve a Transcription Vocabulary
* `ListTranscriptionVocabularies` - List Transcription Vocabularies
* `UpdateTranscriptionVocabulary` - Update a Transcription Vocabulary

### URLSigningKeys

* `CreateURLSigningKey` - Create a URL signing key
* `DeleteURLSigningKey` - Delete a URL signing key
* `GetURLSigningKey` - Retrieve a URL signing key
* `ListURLSigningKeys` - List URL signing keys

### VideoViews

* `GetVideoView` - Get a Video View
* `ListVideoViews` - List Video Views
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
