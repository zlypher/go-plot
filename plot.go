package plot

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// Plotable defines an interface for object which can be represented on a chart.
type Plotable interface {
	// GetX() float64
	GetY() float64
	GetLabel() string
}

// Entry represents a single point on the chart.
type Entry struct {
	Label     string
	LabelAbbr string
	XValue    float64
	YValue    float64
}

// func (e *Entry) GetX() float64 {
// 	return e.XValue
// }

func (e Entry) GetY() float64 {
	return e.YValue
}

func (e Entry) GetLabel() string {
	return e.LabelAbbr
}

// Chart holds all required data to render the chart.
type Chart struct {
	Title   string
	Debug   bool
	Spacing Spacing
	Entries []Plotable
	Theme   Theme
}

// Spacing defines sizes of various spacing elements (margin, padding, ...).
type Spacing struct {
	Margin  int
	Padding int
	Bar     int
	Axis    int
}

// BarChart draws the bar chart to the cmd.
func BarChart(chart Chart) {
	numEntries := len(chart.Entries)
	if numEntries == 0 {
		print(os.Stdout, "No chart entries available")
	}

	width := calculateWidth(chart.Spacing, numEntries)

	if chart.Debug {
		print(os.Stdout, formatDebugInfo(numEntries, width))
	}

	xA := calculateAxis(chart.Entries)
	// TODO: Determine max with of yaxis label
	axisLabelWidth := 5

	if chart.Title != "" {
		print(os.Stdout, formatTitle(chart.Title, width+axisLabelWidth+3))
	}

	print(os.Stdout, formatChart(chart.Entries, xA, chart.Theme, axisLabelWidth))

	print(os.Stdout, formatXAxis(chart.Theme, width, axisLabelWidth))
	print(os.Stdout, formatXAxisLabels(chart.Entries, axisLabelWidth))
	print(os.Stdout, "\n")
}

func calculateAxis(entries []Plotable) Axis {
	low, high, steps := 0.0, math.SmallestNonzeroFloat64, 0.0

	for _, entry := range entries {
		y := entry.GetY()
		if y < low {
			low = y
		} else if y > high {
			high = y
		}
	}

	steps = 1.0

	return Axis{Low: low, High: high, Steps: steps}
}

func formatChart(entries []Plotable, axis Axis, theme Theme, axisLabelWidth int) string {
	// Start the chart with a line with only the y axis drawn
	output := fmt.Sprintf("%s%s\n", strings.Repeat(" ", axisLabelWidth+3), theme.YAxis)

	labelFmt := fmt.Sprintf("%%%d.0f - ", axisLabelWidth)
	for val := axis.High; val >= axis.Low; val -= axis.Steps {
		// Print current y axis value
		output = fmt.Sprintf("%s"+labelFmt, output, val)

		// Print the y axis and the margin until the first bar
		output = fmt.Sprint(output, yAxisChar, strings.Repeat(" ", 2))

		// Print the bars with padding between each bar
		for idx, entry := range entries {
			// If it is not the first element, draw the padding
			if idx != 0 {
				output = fmt.Sprint(output, strings.Repeat(" ", 2))
			}

			// If the bar reaches up to the current value, draw the bar.
			// If not, draw a spacing.
			if entry.GetY() >= val {
				output = fmt.Sprint(output, theme.Bar)
			} else {
				output = fmt.Sprint(output, " ")
			}
		}

		// print remaining margin and newline
		output = fmt.Sprintf("%s%s\n", output, strings.Repeat(" ", 2))
	}

	return output
}

func formatXAxis(theme Theme, width int, axisLabelWidth int) string {
	return fmt.Sprintf("%s%s%s\n",
		strings.Repeat(" ", axisLabelWidth+3),
		theme.CrossAxis,
		strings.Repeat(theme.XAxis, width-1))
}

// TODO: Pass margin and padding as spacing
func formatXAxisLabels(entries []Plotable, axisLabelWidth int) string {
	if len(entries) == 0 {
		return ""
	}

	output := fmt.Sprint(strings.Repeat(" ", axisLabelWidth+6)) // axis + margin

	for idx, entry := range entries {
		if idx != 0 {
			output = fmt.Sprint(output, strings.Repeat(" ", 2)) // pad
		}

		output = fmt.Sprint(output, entry.GetLabel())
	}

	return output
}
