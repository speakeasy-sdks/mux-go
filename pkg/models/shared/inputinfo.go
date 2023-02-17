package shared

type InputInfo struct {
	File     *InputFile     `json:"file,omitempty"`
	Settings *InputSettings `json:"settings,omitempty"`
}
