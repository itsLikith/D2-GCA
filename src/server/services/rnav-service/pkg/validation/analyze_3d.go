package validation

import (
	"fmt"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
)

func ValidateAnalyze3DRequest(
	req *dto.Analyze3DRequest,
) error {

	if len(req.Measurements) < 2 {
		return fmt.Errorf(
			"minimum 2 measurements required",
		)
	}

	for i, m := range req.Measurements {

		if m.SigmaNM <= 0 {
			return fmt.Errorf(
				"invalid sigma at index %d",
				i,
			)
		}
	}

	switch req.AltitudeMode {

	case "RVSM", "FIXED", "CVSM":
		// OK

	default:
		return fmt.Errorf(
			"unknown or unsupported altitude mode: %s",
			req.AltitudeMode,
		)
	}

	return nil
}
