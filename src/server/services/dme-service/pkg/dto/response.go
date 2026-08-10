package dto

type ErrorModelResponse struct {
	RangeNM float64 `json:"rangeNM"`

	SystemErrorNM   float64 `json:"systemErrorNM"`
	AirborneErrorNM float64 `json:"airborneErrorNM"`
	TotalErrorNM    float64 `json:"totalErrorNM"`
}

type PairAccuracyResponse struct {
	InclusionAngleDeg float64 `json:"inclusionAngleDeg"`

	RMSNM      float64 `json:"rmsNM"`
	TwoSigmaNM float64 `json:"twoSigmaNM"`

	RNAV1 bool `json:"rnav1"`
	RNAV2 bool `json:"rnav2"`
}
