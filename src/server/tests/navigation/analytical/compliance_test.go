package analytical_test

import (
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
)

func TestRNAV2Compliant(t *testing.T) {

	// Paper: RNAV-2 allows 2σ ≤ 1.0 NM in paper/compliance.go
	if !analytical.IsRNAV2FlightTechCompliant(0.5) {
		t.Error("0.5 should be RNAV-2 compliant")
	}

	if !analytical.IsRNAV2FlightTechCompliant(1.0) {
		t.Error("1.0 should be RNAV-2 compliant")
	}

	if analytical.IsRNAV2FlightTechCompliant(1.1) {
		t.Error("1.1 should NOT be RNAV-2 compliant")
	}
}

func TestRNAV1Compliant(t *testing.T) {

	// Paper: RNAV-1 allows 2σ ≤ 0.5 NM in paper/compliance.go
	if !analytical.IsRNAV1FlightTechCompliant(0.3) {
		t.Error("0.3 should be RNAV-1 compliant")
	}

	if !analytical.IsRNAV1FlightTechCompliant(0.5) {
		t.Error("0.5 should be RNAV-1 compliant")
	}

	if analytical.IsRNAV1FlightTechCompliant(0.6) {
		t.Error("0.6 should NOT be RNAV-1 compliant")
	}
}
