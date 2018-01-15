package plot

import (
	"testing"
)

func TestEntryStruct(t *testing.T) {
	tests := []struct {
		entry Entry
		y     float64
		label string
	}{
		{Entry{LabelAbbr: "Test", YValue: 1}, 1, "Test"},
	}

	for _, test := range tests {
		y := test.entry.GetY()
		if y != test.y {
			t.Errorf("GetY(%v) was incorrect, got: \"%f\", want: \"%f\".", test.entry, y, test.y)
		}

		label := test.entry.GetLabel()
		if label != test.label {
			t.Errorf("GetLabel(%v) was incorrect, got: \"%s\", want: \"%s\".", test.entry, label, test.label)
		}
	}
}

func TestFormatXAxis(t *testing.T) {
	tests := []struct {
		theme      Theme
		width      int
		labelWidth int
		formatted  string
	}{
		{Theme{CrossAxis: "+", XAxis: "-"}, 3, 3, "      +--\n"},
		{Theme{CrossAxis: "+", XAxis: "-"}, 10, 3, "      +---------\n"},
		{Theme{CrossAxis: "+", XAxis: "+"}, 10, 3, "      ++++++++++\n"},
		{Theme{CrossAxis: "#", XAxis: "="}, 10, 3, "      #=========\n"},
	}

	for _, test := range tests {
		formatted := formatXAxis(test.theme, test.width, test.labelWidth)
		if formatted != test.formatted {
			t.Errorf("formatXAxis(%v, %d, %d) was incorrect, got: \"%s\", want: \"%s\".", test.theme, test.width, test.labelWidth, formatted, test.formatted)
		}
	}
}

func TestFormatXAxisLabels_NoEntries(t *testing.T) {
	var entries []Plotable
	labelWidth := 4
	expected := ""

	actual := formatXAxisLabels(entries, labelWidth)

	if actual != expected {
		t.Errorf("formatXAxisLabels(..., %d) was incorrect, got: \"%s\", want \"%s\".", labelWidth, actual, expected)
	}
}

func TestFormatXAxisLabels_OneEntry(t *testing.T) {
	var entries []Plotable
	entries = append(entries, Entry{LabelAbbr: "A"})
	labelWidth := 4
	expected := "          A"

	actual := formatXAxisLabels(entries, labelWidth)

	if actual != expected {
		t.Errorf("formatXAxisLabels(..., %d) was incorrect, got: \"%s\", want \"%s\".", labelWidth, actual, expected)
	}
}

func TestFormatXAxisLabels_ManyEntries(t *testing.T) {
	var entries []Plotable
	entries = append(entries, Entry{LabelAbbr: "A"})
	entries = append(entries, Entry{LabelAbbr: "B"})
	entries = append(entries, Entry{LabelAbbr: "C"})
	entries = append(entries, Entry{LabelAbbr: "D"})
	labelWidth := 3
	expected := "         A  B  C  D"

	actual := formatXAxisLabels(entries, labelWidth)

	if actual != expected {
		t.Errorf("formatXAxisLabels(..., %d) was incorrect, got: \"%s\", want \"%s\".", labelWidth, actual, expected)
	}
}
