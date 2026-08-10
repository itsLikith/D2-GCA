package dto

type Measurement struct {
	AzimuthDeg   float64 `json:"azimuthDeg"`
	ElevationDeg float64 `json:"elevationDeg"`
	SigmaNM      float64 `json:"sigmaNM"`
}

type AnalyzeRequest struct {
	Measurements []Measurement `json:"measurements"`

	Observations []float64 `json:"observations"`
}

type Analyze3DRequest struct {
	Measurements []Measurement `json:"measurements"`

	// AltitudeMode: "RVSM", "FIXED", "CVSM"
	AltitudeMode string `json:"altitudeMode"`
}
