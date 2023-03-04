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

type playbackRestrictions struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newPlaybackRestrictions(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *playbackRestrictions {
	return &playbackRestrictions{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreatePlaybackRestriction - Create a Playback Restriction
// Create a new Playback Restriction.
func (s *playbackRestrictions) CreatePlaybackRestriction(ctx context.Context, request operations.CreatePlaybackRestrictionRequest) (*operations.CreatePlaybackRestrictionResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/video/v1/playback-restrictions"

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

	retryConfig := request.Retries
	if retryConfig == nil {
		retryConfig = &utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 500,
				MaxInterval:     60000,
				Exponent:        1.5,
				MaxElapsedTime:  3600000,
			},
			RetryConnectionErrors: true,
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreatePlaybackRestrictionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 201:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PlaybackRestrictionResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PlaybackRestrictionResponse = out
		}
	}

	return res, nil
}

// DeletePlaybackRestriction - Delete a Playback Restriction
// Deletes a single Playback Restriction.
func (s *playbackRestrictions) DeletePlaybackRestriction(ctx context.Context, request operations.DeletePlaybackRestrictionRequest) (*operations.DeletePlaybackRestrictionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/playback-restrictions/{PLAYBACK_RESTRICTION_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	retryConfig := request.Retries
	if retryConfig == nil {
		retryConfig = &utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 500,
				MaxInterval:     60000,
				Exponent:        1.5,
				MaxElapsedTime:  3600000,
			},
			RetryConnectionErrors: true,
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeletePlaybackRestrictionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 204:
	}

	return res, nil
}

// GetPlaybackRestriction - Retrieve a Playback Restriction
// Retrieves a Playback Restriction associated with the unique identifier.
func (s *playbackRestrictions) GetPlaybackRestriction(ctx context.Context, request operations.GetPlaybackRestrictionRequest) (*operations.GetPlaybackRestrictionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/playback-restrictions/{PLAYBACK_RESTRICTION_ID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	retryConfig := request.Retries
	if retryConfig == nil {
		retryConfig = &utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 500,
				MaxInterval:     60000,
				Exponent:        1.5,
				MaxElapsedTime:  3600000,
			},
			RetryConnectionErrors: true,
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetPlaybackRestrictionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PlaybackRestrictionResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PlaybackRestrictionResponse = out
		}
	}

	return res, nil
}

// ListPlaybackRestrictions - List Playback Restrictions
// Returns a list of all Playback Restrictions.
func (s *playbackRestrictions) ListPlaybackRestrictions(ctx context.Context, request operations.ListPlaybackRestrictionsRequest) (*operations.ListPlaybackRestrictionsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/video/v1/playback-restrictions"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if err := utils.PopulateQueryParams(ctx, req, request.QueryParams); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.securityClient

	retryConfig := request.Retries
	if retryConfig == nil {
		retryConfig = &utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 500,
				MaxInterval:     60000,
				Exponent:        1.5,
				MaxElapsedTime:  3600000,
			},
			RetryConnectionErrors: true,
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListPlaybackRestrictionsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ListPlaybackRestrictionsResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListPlaybackRestrictionsResponse = out
		}
	}

	return res, nil
}

// UpdateReferrerDomainRestriction - Update the Referrer Playback Restriction
// Allows you to modify the list of domains or change how Mux validates playback requests without the `Referer` HTTP header. The Referrer restriction fully replaces the old list with this new list of domains.
func (s *playbackRestrictions) UpdateReferrerDomainRestriction(ctx context.Context, request operations.UpdateReferrerDomainRestrictionRequest) (*operations.UpdateReferrerDomainRestrictionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/video/v1/playback-restrictions/{PLAYBACK_RESTRICTION_ID}/referrer", request.PathParams)

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

	retryConfig := request.Retries
	if retryConfig == nil {
		retryConfig = &utils.RetryConfig{
			Strategy: "backoff",
			Backoff: &utils.BackoffStrategy{
				InitialInterval: 500,
				MaxInterval:     60000,
				Exponent:        1.5,
				MaxElapsedTime:  3600000,
			},
			RetryConnectionErrors: true,
		}
	}

	httpRes, err := utils.Retry(ctx, utils.Retries{
		Config: retryConfig,
		StatusCodes: []string{
			"5XX",
		},
	}, func() (*http.Response, error) {
		return client.Do(req)
	})
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateReferrerDomainRestrictionResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.PlaybackRestrictionResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.PlaybackRestrictionResponse = out
		}
	}

	return res, nil
}
