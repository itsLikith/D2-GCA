package dto

type ErrorModelRequest struct {
	RangeNM float64 `json:"rangeNM"`
}

type PairAccuracyRequest struct {
	Range1NM float64 `json:"range1NM"`
	Range2NM float64 `json:"range2NM"`

	InclusionAngleDeg float64 `json:"inclusionAngleDeg"`
}
