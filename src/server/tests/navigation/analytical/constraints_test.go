package analytical_test

import (
	"testing"

	"github.com/D2-GCA/PRISM/pkg/navigation/analytical"
)

func TestValidInclusionAngle(t *testing.T) {

	cases := []struct {
		angle float64
		valid bool
	}{
		{29.0, false},
		{30.0, true},
		{90.0, true},
		{150.0, true},
		{151.0, false},
	}

	for _, tc := range cases {

		result := analytical.IsValidInclusionAngle(tc.angle)

		if result != tc.valid {
			t.Errorf("angle=%.0f: expected %v, got %v",
				tc.angle, tc.valid, result)
		}
	}
}

func TestPreferredRNAV1Angle(t *testing.T) {

	cases := []struct {
		angle     float64
		preferred bool
	}{
		{39.0, false},
		{40.0, true},
		{90.0, true},
		{140.0, true},
		{141.0, false},
	}

	for _, tc := range cases {

		result := analytical.IsPreferredRNAV1Angle(tc.angle)

		if result != tc.preferred {
			t.Errorf("angle=%.0f: expected %v, got %v",
				tc.angle, tc.preferred, result)
		}
	}
}
