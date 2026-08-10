package validation

import (
	"fmt"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
)

func ValidatePairAccuracyRequest(
	req *dto.PairAccuracyRequest,
) error {

	if req.Range1NM <= 0 || req.Range2NM <= 0 {
		return fmt.Errorf(
			"ranges must be positive",
		)
	}

	if req.InclusionAngleDeg <= 0 ||
		req.InclusionAngleDeg >= 180 {

		return fmt.Errorf(
			"inclusion angle must be between 0° and 180° (exclusive)",
		)
	}

	return nil
}
