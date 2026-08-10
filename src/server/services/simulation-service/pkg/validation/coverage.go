package validation

import (
	"fmt"

	"github.com/D2-GCA/PRISM/simulation-service/pkg/dto"
)

func ValidateCoverageRequest(
	req *dto.CoverageRequest,
) error {

	if len(req.Stations) < 2 {
		return fmt.Errorf(
			"minimum 2 stations required",
		)
	}

	if req.GridStepNM <= 0 {
		return fmt.Errorf(
			"gridStepNM must be positive",
		)
	}

	if req.MinX >= req.MaxX {
		return fmt.Errorf(
			"minX must be less than maxX",
		)
	}

	if req.MinY >= req.MaxY {
		return fmt.Errorf(
			"minY must be less than maxY",
		)
	}

	return nil
}
