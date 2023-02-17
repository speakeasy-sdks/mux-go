package mux

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/mux-go/pkg/models/operations"
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
	"strings"
)

type liveStreams struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newLiveStreams(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *liveStreams {
	return &liveStreams{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateLiveStream - Create a live stream
// Creates a new live stream. Once created, an encoder can connect to Mux via the specified stream key and begin streaming to an audience.
func (s *liveStreams) CreateLiveStream(ctx context.Context, request operations.CreateLiveStreamRequest) (*operations.CreateLiveStreamResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/video/v1/live-streams"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateLiveStreamResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LiveStreamResponse = out
		}
	}

	return res, nil
}

// CreateLiveStreamPlaybackID - Create a live stream playback ID
// Create a new playback ID for this live stream, through which a viewer can watch the streamed content of the live stream.
func (s *liveStreams) CreateLiveStreamPlaybackID(ctx context.Context, request operations.CreateLiveStreamPlaybackIDRequest) (*operations.CreateLiveStreamPlaybackIDResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateLiveStreamPlaybackIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.CreatePlaybackIDResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreatePlaybackIDResponse = out
		}
	}

	return res, nil
}

// CreateLiveStreamSimulcastTarget - Create a live stream simulcast target
// Create a simulcast target for the parent live stream. Simulcast target can only be created when the parent live stream is in idle state. Only one simulcast target can be created at a time with this API.
func (s *liveStreams) CreateLiveStreamSimulcastTarget(ctx context.Context, request operations.CreateLiveStreamSimulcastTargetRequest) (*operations.CreateLiveStreamSimulcastTargetResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateLiveStreamSimulcastTargetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SimulcastTargetResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SimulcastTargetResponse = out
		}
	}

	return res, nil
}

// DeleteLiveStream - Delete a live stream
// Deletes a live stream from the current environment. If the live stream is currently active and being streamed to, ingest will be terminated and the encoder will be disconnected.
func (s *liveStreams) DeleteLiveStream(ctx context.Context, request operations.DeleteLiveStreamRequest) (*operations.DeleteLiveStreamResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteLiveStreamResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// DeleteLiveStreamPlaybackID - Delete a live stream playback ID
// Deletes the playback ID for the live stream. This will not disable ingest (as the live stream still exists). New attempts to play back the live stream will fail immediately. However, current viewers will be able to continue watching the stream for some period of time.
func (s *liveStreams) DeleteLiveStreamPlaybackID(ctx context.Context, request operations.DeleteLiveStreamPlaybackIDRequest) (*operations.DeleteLiveStreamPlaybackIDResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteLiveStreamPlaybackIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// DeleteLiveStreamSimulcastTarget - Delete a Live Stream Simulcast Target
// Delete the simulcast target using the simulcast target ID returned when creating the simulcast target. Simulcast Target can only be deleted when the parent live stream is in idle state.
func (s *liveStreams) DeleteLiveStreamSimulcastTarget(ctx context.Context, request operations.DeleteLiveStreamSimulcastTargetRequest) (*operations.DeleteLiveStreamSimulcastTargetResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteLiveStreamSimulcastTargetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// DisableLiveStream - Disable a live stream
// Disables a live stream, making it reject incoming RTMP streams until re-enabled. The API also ends the live stream recording immediately when active. Ending the live stream recording adds the `EXT-X-ENDLIST` tag to the HLS manifest which notifies the player that this live stream is over.
//
// Mux also closes the encoder connection immediately. Any attempt from the encoder to re-establish connection will fail till the live stream is re-enabled.
func (s *liveStreams) DisableLiveStream(ctx context.Context, request operations.DisableLiveStreamRequest) (*operations.DisableLiveStreamResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/disable", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DisableLiveStreamResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.DisableLiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.DisableLiveStreamResponse = out
		}
	}

	return res, nil
}

// EnableLiveStream - Enable a live stream
// Enables a live stream, allowing it to accept an incoming RTMP stream.
func (s *liveStreams) EnableLiveStream(ctx context.Context, request operations.EnableLiveStreamRequest) (*operations.EnableLiveStreamResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/enable", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.EnableLiveStreamResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EnableLiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.EnableLiveStreamResponse = out
		}
	}

	return res, nil
}

// GetLiveStream - Retrieve a live stream
// Retrieves the details of a live stream that has previously been created. Supply the unique live stream ID that was returned from your previous request, and Mux will return the corresponding live stream information. The same information is returned when creating a live stream.
func (s *liveStreams) GetLiveStream(ctx context.Context, request operations.GetLiveStreamRequest) (*operations.GetLiveStreamResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetLiveStreamResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LiveStreamResponse = out
		}
	}

	return res, nil
}

// GetLiveStreamPlaybackID - Retrieve a live stream playback ID
// Fetches information about a live stream's playback ID, through which a viewer can watch the streamed content from this live stream.
func (s *liveStreams) GetLiveStreamPlaybackID(ctx context.Context, request operations.GetLiveStreamPlaybackIDRequest) (*operations.GetLiveStreamPlaybackIDResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetLiveStreamPlaybackIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.GetLiveStreamPlaybackIDResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetLiveStreamPlaybackIDResponse = out
		}
	}

	return res, nil
}

// GetLiveStreamSimulcastTarget - Retrieve a Live Stream Simulcast Target
// Retrieves the details of the simulcast target created for the parent live stream. Supply the unique live stream ID and simulcast target ID that was returned in the response of create simulcast target request, and Mux will return the corresponding information.
func (s *liveStreams) GetLiveStreamSimulcastTarget(ctx context.Context, request operations.GetLiveStreamSimulcastTargetRequest) (*operations.GetLiveStreamSimulcastTargetResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetLiveStreamSimulcastTargetResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SimulcastTargetResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SimulcastTargetResponse = out
		}
	}

	return res, nil
}

// ListLiveStreams - List live streams
// Lists the live streams that currently exist in the current environment.
func (s *liveStreams) ListLiveStreams(ctx context.Context, request operations.ListLiveStreamsRequest) (*operations.ListLiveStreamsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/video/v1/live-streams"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListLiveStreamsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ListLiveStreamsResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListLiveStreamsResponse = out
		}
	}

	return res, nil
}

// ResetStreamKey - Reset a live stream's stream key
// Reset a live stream key if you want to immediately stop the current stream key from working and create a new stream key that can be used for future broadcasts.
func (s *liveStreams) ResetStreamKey(ctx context.Context, request operations.ResetStreamKeyRequest) (*operations.ResetStreamKeyResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/reset-stream-key", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ResetStreamKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LiveStreamResponse = out
		}
	}

	return res, nil
}

// SignalLiveStreamComplete - Signal a live stream is finished
// (Optional) End the live stream recording immediately instead of waiting for the reconnect_window. `EXT-X-ENDLIST` tag is added to the HLS manifest which notifies the player that this live stream is over.
//
// Mux does not close the encoder connection immediately. Encoders are often configured to re-establish connections immediately which would result in a new recorded asset. For this reason, Mux waits for 60s before closing the connection with the encoder. This 60s timeframe is meant to give encoder operators a chance to disconnect from their end.
func (s *liveStreams) SignalLiveStreamComplete(ctx context.Context, request operations.SignalLiveStreamCompleteRequest) (*operations.SignalLiveStreamCompleteResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/complete", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.SignalLiveStreamCompleteResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SignalLiveStreamCompleteResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.SignalLiveStreamCompleteResponse = out
		}
	}

	return res, nil
}

// UpdateLiveStream - Update a live stream
// Updates the parameters of a previously-created live stream. This currently supports a subset of variables. Supply the live stream ID and the updated parameters and Mux will return the corresponding live stream information. The information returned will be the same after update as for subsequent get live stream requests.
func (s *liveStreams) UpdateLiveStream(ctx context.Context, request operations.UpdateLiveStreamRequest) (*operations.UpdateLiveStreamResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateLiveStreamResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LiveStreamResponse = out
		}
	}

	return res, nil
}

// UpdateLiveStreamEmbeddedSubtitles - Update a live stream's embedded subtitles
// Configures a live stream to receive embedded closed captions.
// The resulting Asset's subtitle text track will have `closed_captions: true` set.
func (s *liveStreams) UpdateLiveStreamEmbeddedSubtitles(ctx context.Context, request operations.UpdateLiveStreamEmbeddedSubtitlesRequest) (*operations.UpdateLiveStreamEmbeddedSubtitlesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/embedded-subtitles", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateLiveStreamEmbeddedSubtitlesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LiveStreamResponse = out
		}
	}

	return res, nil
}

// UpdateLiveStreamGeneratedSubtitles - Update a live stream's generated subtitles
// Updates a live stream's automatic-speech-recognition-generated subtitle configuration.
// Automatic speech recognition subtitles can be removed by sending an empty array in the
// request payload.
func (s *liveStreams) UpdateLiveStreamGeneratedSubtitles(ctx context.Context, request operations.UpdateLiveStreamGeneratedSubtitlesRequest) (*operations.UpdateLiveStreamGeneratedSubtitlesResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/live-streams/{LIVE_STREAM_ID}/generated-subtitles", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateLiveStreamGeneratedSubtitlesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.LiveStreamResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.LiveStreamResponse = out
		}
	}

	return res, nil
}
