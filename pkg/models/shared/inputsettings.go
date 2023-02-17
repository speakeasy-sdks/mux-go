package shared

type InputSettingsOverlaySettingsHorizontalAlignEnum string

const (
	InputSettingsOverlaySettingsHorizontalAlignEnumLeft   InputSettingsOverlaySettingsHorizontalAlignEnum = "left"
	InputSettingsOverlaySettingsHorizontalAlignEnumCenter InputSettingsOverlaySettingsHorizontalAlignEnum = "center"
	InputSettingsOverlaySettingsHorizontalAlignEnumRight  InputSettingsOverlaySettingsHorizontalAlignEnum = "right"
)

type InputSettingsOverlaySettingsVerticalAlignEnum string

const (
	InputSettingsOverlaySettingsVerticalAlignEnumTop    InputSettingsOverlaySettingsVerticalAlignEnum = "top"
	InputSettingsOverlaySettingsVerticalAlignEnumMiddle InputSettingsOverlaySettingsVerticalAlignEnum = "middle"
	InputSettingsOverlaySettingsVerticalAlignEnumBottom InputSettingsOverlaySettingsVerticalAlignEnum = "bottom"
)

// InputSettingsOverlaySettings
// An object that describes how the image file referenced in URL should be placed over the video (i.e. watermarking). Ensure that the URL is active and persists the entire lifespan of the video object.
type InputSettingsOverlaySettings struct {
	Height           *string                                          `json:"height,omitempty"`
	HorizontalAlign  *InputSettingsOverlaySettingsHorizontalAlignEnum `json:"horizontal_align,omitempty"`
	HorizontalMargin *string                                          `json:"horizontal_margin,omitempty"`
	Opacity          *string                                          `json:"opacity,omitempty"`
	VerticalAlign    *InputSettingsOverlaySettingsVerticalAlignEnum   `json:"vertical_align,omitempty"`
	VerticalMargin   *string                                          `json:"vertical_margin,omitempty"`
	Width            *string                                          `json:"width,omitempty"`
}

type InputSettingsTextTypeEnum string

const (
	InputSettingsTextTypeEnumSubtitles InputSettingsTextTypeEnum = "subtitles"
)

type InputSettingsTypeEnum string

const (
	InputSettingsTypeEnumVideo InputSettingsTypeEnum = "video"
	InputSettingsTypeEnumAudio InputSettingsTypeEnum = "audio"
	InputSettingsTypeEnumText  InputSettingsTypeEnum = "text"
)

// InputSettings
// An array of objects that each describe an input file to be used to create the asset. As a shortcut, `input` can also be a string URL for a file when only one input file is used. See `input[].url` for requirements.
type InputSettings struct {
	ClosedCaptions  *bool                         `json:"closed_captions,omitempty"`
	EndTime         *float64                      `json:"end_time,omitempty"`
	LanguageCode    *string                       `json:"language_code,omitempty"`
	Name            *string                       `json:"name,omitempty"`
	OverlaySettings *InputSettingsOverlaySettings `json:"overlay_settings,omitempty"`
	Passthrough     *string                       `json:"passthrough,omitempty"`
	StartTime       *float64                      `json:"start_time,omitempty"`
	TextType        *InputSettingsTextTypeEnum    `json:"text_type,omitempty"`
	Type            *InputSettingsTypeEnum        `json:"type,omitempty"`
	URL             *string                       `json:"url,omitempty"`
}
