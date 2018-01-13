package plot

import (
	"testing"
)

func TestPrintDebugInfo(t *testing.T) {
	// TODO: Make printDebugInfo testable
}

func TestPrintTitle(t *testing.T) {
	// TODO: Make printTitle testable
}

func TestCalculateWidth(t *testing.T) {
	tests := []struct {
		num     int
		spacing Spacing
		width   int
	}{
		{0, Spacing{Margin: 2, Padding: 2, Axis: 1, Bar: 1}, 3},
		{2, Spacing{Margin: 2, Padding: 2, Axis: 1, Bar: 1}, 9},
		{4, Spacing{Margin: 2, Padding: 2, Axis: 1, Bar: 1}, 15},
		{2, Spacing{Margin: 2, Padding: 2, Axis: 1, Bar: 2}, 11},
		{2, Spacing{Margin: 4, Padding: 2, Axis: 1, Bar: 2}, 15},
		{2, Spacing{Margin: 4, Padding: 4, Axis: 1, Bar: 2}, 17},
	}

	for _, test := range tests {
		width := calculateWidth(test.spacing, test.num)
		if width != test.width {
			t.Errorf("Calculated width (%d, %v) was incorrect, got: %d, want: %d.", test.num, test.spacing, width, test.width)
		}
	}
}
