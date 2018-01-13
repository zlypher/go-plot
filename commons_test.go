package plot

import (
	"bytes"
	"testing"
)

func TestPrintDebugInfo(t *testing.T) {
	// TODO: Make printDebugInfo testable
}

func TestFormatTitle(t *testing.T) {
	tests := []struct {
		title     string
		width     int
		formatted string
	}{
		{"", 0, "\n\n"},
		{"A", 1, "A\n\n"},
		{"ABC", 3, "ABC\n\n"},
		{"A", 5, "  A\n\n"},
		{"Hello", 12, "   Hello\n\n"},
	}

	for _, test := range tests {
		formatted := formatTitle(test.title, test.width)
		if formatted != test.formatted {
			t.Errorf("FormatTitle(%s, %d) was incorrect, got: \"%s\", want: \"%s\".", test.title, test.width, formatted, test.formatted)
		}
	}
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

func TestPrint(t *testing.T) {
	tests := []struct {
		output string
	}{
		{"A"},
		{"Hello World"},
		{"This is a Test"},
		{"Lorem Ipsum"},
	}

	for _, test := range tests {
		var b bytes.Buffer
		print(&b, test.output)
		if b.String() != test.output {
			t.Errorf("print(%s) was incorrect, got: %s, want: %s.", test.output, b.String(), test.output)
		}
	}
}
