package dto

type AnalyzeResponse struct {
	X float64 `json:"x"`

	Y float64 `json:"y"`

	RMSAccuracyNM float64 `json:"rmsAccuracyNM"`

	TwoSigmaNM float64 `json:"twoSigmaNM"`

	RNAV1 bool `json:"rnav1"`

	RNAV2 bool `json:"rnav2"`
}

type Analyze3DResponse struct {
	HorizontalRMSNM float64 `json:"horizontalRmsNM"`

	VerticalRMSNM float64 `json:"verticalRmsNM"`

	TwoSigmaNM float64 `json:"twoSigmaNM"`

	AltitudeMode string `json:"altitudeMode"`

	RNAV1 bool `json:"rnav1"`

	RNAV2 bool `json:"rnav2"`
}
