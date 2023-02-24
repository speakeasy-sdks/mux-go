<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
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
            },
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

    ctx := context.Background()
    res, err := s.Assets.CreateAsset(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.AssetResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->