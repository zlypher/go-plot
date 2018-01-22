package plot

import (
	"testing"
)

func TestGetExtremes(t *testing.T) {
	var entries []Plotable
	entries = append(entries, Entry{YValue: 10})
	entries = append(entries, Entry{YValue: 30})
	entries = append(entries, Entry{YValue: -10})
	entries = append(entries, Entry{YValue: 30})
	entries = append(entries, Entry{YValue: 0})
	entries = append(entries, Entry{YValue: 20})
	entries = append(entries, Entry{YValue: 19})
	expMin, expMax := -10.0, 30.0
	min, max := getExtremes(entries)

	if min != expMin {
		t.Errorf("getExtremes(%v), min was incorrect, got: \"%f\", want: \"%f\".", entries, min, expMin)
	}

	if max != expMax {
		t.Errorf("getExtremes(%v), max was incorrect, got: \"%f\", want: \"%f\".", entries, max, expMax)
	}
}

func TestGetExtremes_OneEntry(t *testing.T) {
	var entries []Plotable
	entries = append(entries, Entry{YValue: 10})
	expMin, expMax := 10.0, 10.0
	min, max := getExtremes(entries)

	if min != expMin {
		t.Errorf("getExtremes(%v), min was incorrect, got: \"%f\", want: \"%f\".", entries, min, expMin)
	}

	if max != expMax {
		t.Errorf("getExtremes(%v), max was incorrect, got: \"%f\", want: \"%f\".", entries, max, expMax)
	}
}

// TODO
func TestCalculateAxis(t *testing.T) {
	expected := Axis{High: 0.0, Steps: 1.0}
	actual := calculateAxis(nil)

	if actual != expected {
		t.Errorf("calculateAxis(%v) was incorrect, got: \"%v\", want: \"%v\".", nil, actual, expected)
	}
}

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

func TestFormatChart_BasicImplementation(t *testing.T) {
	var entries []Plotable
	entries = append(entries, Entry{XValue: 1, YValue: 1})
	entries = append(entries, Entry{XValue: 2, YValue: 3})
	entries = append(entries, Entry{XValue: 3, YValue: 2})
	axis := Axis{Low: 0, Steps: 1, High: 5}
	theme := Theme{YAxis: "|", Bar: "+"}
	axisLabelWidth := 5
	expected := "        |\n    5 - |           \n    4 - |           \n    3 - |     +     \n    2 - |     +  +  \n    1 - |  +  +  +  \n"
	formatted := formatChart(entries, axis, theme, axisLabelWidth)

	if formatted != expected {
		t.Errorf("formatChart(%v, %v, %v, %d) was incorrect, got: \"%s\", want: \"%s\".", entries, axis, theme, axisLabelWidth, formatted, expected)
	}
}

func TestFormatXAxis(t *testing.T) {
	tests := []struct {
		theme      Theme
		width      int
		labelWidth int
		axisVal    float64
		formatted  string
	}{
		{Theme{CrossAxis: "+", XAxis: "-"}, 3, 3, 0.0, "  0 - +--\n"},
		{Theme{CrossAxis: "+", XAxis: "-"}, 10, 3, -2.0, " -2 - +---------\n"},
		{Theme{CrossAxis: "+", XAxis: "+"}, 10, 3, 3.0, "  3 - ++++++++++\n"},
		{Theme{CrossAxis: "#", XAxis: "="}, 10, 3, 5.0, "  5 - #=========\n"},
	}

	for _, test := range tests {
		formatted := formatXAxis(test.theme, test.width, test.labelWidth, test.axisVal)
		if formatted != test.formatted {
			t.Errorf("formatXAxis(%v, %d, %d, %f) was incorrect, got: \"%s\", want: \"%s\".", test.theme, test.width, test.labelWidth, test.axisVal, formatted, test.formatted)
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
