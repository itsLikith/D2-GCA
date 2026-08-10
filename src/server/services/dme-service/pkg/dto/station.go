package dto

type StationDTO struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	X               float64 `json:"x"`
	Y               float64 `json:"y"`
	ElevationFt     float64 `json:"elevationFt"`
	ServiceVolumeNM float64 `json:"serviceVolumeNM"`
}
