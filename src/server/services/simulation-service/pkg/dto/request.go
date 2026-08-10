package dto

type StationDTO struct {
	ID              string  `json:"id"`
	X               float64 `json:"x"`
	Y               float64 `json:"y"`
	ElevationFt     float64 `json:"elevationFt"`
	ServiceVolumeNM float64 `json:"serviceVolumeNM"`
}

type CoverageRequest struct {
	Stations []StationDTO `json:"stations"`

	MinX float64 `json:"minX"`
	MaxX float64 `json:"maxX"`
	MinY float64 `json:"minY"`
	MaxY float64 `json:"maxY"`

	GridStepNM float64 `json:"gridStepNM"`
}

type ElevationSweepRequest struct {
	Sigma1NM float64 `json:"sigma1NM"`
	Sigma2NM float64 `json:"sigma2NM"`

	// AltitudeMode: "RVSM", "FIXED", "CVSM"
	AltitudeMode string `json:"altitudeMode"`

	InclusionAnglesDeg []float64 `json:"inclusionAnglesDeg"`

	ElevationMinDeg  float64 `json:"elevationMinDeg"`
	ElevationMaxDeg  float64 `json:"elevationMaxDeg"`
	ElevationStepDeg float64 `json:"elevationStepDeg"`
}
