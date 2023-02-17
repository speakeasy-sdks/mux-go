package shared

type IncidentStatusEnum string

const (
	IncidentStatusEnumOpen    IncidentStatusEnum = "open"
	IncidentStatusEnumClosed  IncidentStatusEnum = "closed"
	IncidentStatusEnumExpired IncidentStatusEnum = "expired"
)
