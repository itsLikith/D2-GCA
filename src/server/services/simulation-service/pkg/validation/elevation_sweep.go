package validation

import (
	"fmt"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
)

func ValidateElevationSweepRequest(
	req *dto.ElevationSweepRequest,
) error {

	if req.Sigma1NM <= 0 || req.Sigma2NM <= 0 {
		return fmt.Errorf(
			"sigma values must be positive",
		)
	}

	if len(req.InclusionAnglesDeg) == 0 {
		return fmt.Errorf(
			"at least one inclusion angle required",
		)
	}

	if req.ElevationStepDeg <= 0 {
		return fmt.Errorf(
			"elevation step must be positive",
		)
	}

	mode := req.AltitudeMode

	if mode != "RVSM" &&
		mode != "FIXED" &&
		mode != "CVSM" {

		return fmt.Errorf(
			"invalid altitude mode: %s (must be RVSM, FIXED, or CVSM)",
			mode,
		)
	}

	return nil
}
