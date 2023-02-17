package shared

type LiveStreamStatusEnum string

const (
	LiveStreamStatusEnumActive   LiveStreamStatusEnum = "active"
	LiveStreamStatusEnumIdle     LiveStreamStatusEnum = "idle"
	LiveStreamStatusEnumDisabled LiveStreamStatusEnum = "disabled"
)
