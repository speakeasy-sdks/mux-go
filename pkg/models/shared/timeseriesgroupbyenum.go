package shared

type TimeseriesGroupByEnum string

const (
	TimeseriesGroupByEnumTenMinutes TimeseriesGroupByEnum = "ten_minutes"
	TimeseriesGroupByEnumHour       TimeseriesGroupByEnum = "hour"
	TimeseriesGroupByEnumDay        TimeseriesGroupByEnum = "day"
)
