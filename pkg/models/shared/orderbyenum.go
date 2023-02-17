package shared

type OrderByEnum string

const (
	OrderByEnumNegativeImpact OrderByEnum = "negative_impact"
	OrderByEnumValue          OrderByEnum = "value"
	OrderByEnumViews          OrderByEnum = "views"
	OrderByEnumField          OrderByEnum = "field"
)
