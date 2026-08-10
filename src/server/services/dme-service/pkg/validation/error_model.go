package validation

import (
	"fmt"

	"github.com/D2-GCA/PRISM/dme-service/pkg/dto"
)

func ValidateErrorModelRequest(
	req *dto.ErrorModelRequest,
) error {

	if req.RangeNM <= 0 {
		return fmt.Errorf(
			"rangeNM must be positive",
		)
	}

	if req.RangeNM > 200 {
		return fmt.Errorf(
			"rangeNM exceeds maximum DME range",
		)
	}

	return nil
}
