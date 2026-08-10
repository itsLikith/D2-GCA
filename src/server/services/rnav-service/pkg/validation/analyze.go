package validation

import (
	"fmt"

	"github.com/D2-GCA/PRISM/rnav-service/pkg/dto"
)

func ValidateAnalyzeRequest(
	req *dto.AnalyzeRequest,
) error {

	if len(req.Measurements) < 2 {
		return fmt.Errorf(
			"minimum 2 measurements required",
		)
	}

	if len(req.Observations) != len(req.Measurements) {
		return fmt.Errorf(
			"observations count mismatch",
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

	return nil
}
