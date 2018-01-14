package plot

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/zlypher/go-plot/chart"
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

	printChart(chart.Entries, xA, chart.Theme, axisLabelWidth)

	print(os.Stdout, formatXAxis(chart.Theme, width, axisLabelWidth))
	printXAxisLabels(chart.Entries, axisLabelWidth)
	print(os.Stdout, "\n")
}

func calculateAxis(entries []Plotable) chart.Axis {
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

	return chart.Axis{Low: low, High: high, Steps: steps}
}

func printChart(entries []Plotable, axis chart.Axis, theme Theme, axisLabelWidth int) {
	// Start the chart with a line with only the y axis drawn
	fmt.Printf("%s%s\n", strings.Repeat(" ", axisLabelWidth+3), theme.YAxis)

	labelFmt := fmt.Sprintf("%%%d.0f - ", axisLabelWidth)
	for val := axis.High; val >= axis.Low; val -= axis.Steps {
		// Print current y axis value
		fmt.Printf(labelFmt, val)

		// Print the y axis and the margin until the first bar
		fmt.Printf("%s%s", yAxisChar, strings.Repeat(" ", 2))

		// Print the bars with padding between each bar
		for idx, entry := range entries {
			// If it is not the first element, draw the padding
			if idx != 0 {
				fmt.Printf(strings.Repeat(" ", 2))
			}

			// If the bar reaches up to the current value, draw the bar.
			// If not, draw a spacing.
			if entry.GetY() >= val {
				fmt.Printf(theme.Bar)
			} else {
				fmt.Printf(" ")
			}
		}

		// print remaining margin and newline
		fmt.Printf("%s\n", strings.Repeat(" ", 2))
	}
}

func formatXAxis(theme Theme, width int, axisLabelWidth int) string {
	return fmt.Sprintf("%s%s%s\n",
		strings.Repeat(" ", axisLabelWidth+3),
		theme.CrossAxis,
		strings.Repeat(theme.XAxis, width-1))
}

func printXAxisLabels(entries []Plotable, axisLabelWidth int) {
	fmt.Printf("%s%s",
		strings.Repeat(" ", axisLabelWidth+3),
		strings.Repeat(" ", 3)) // axis + margin

	for idx, entry := range entries {
		if idx != 0 {
			fmt.Printf(strings.Repeat(" ", 2)) // pad
		}

		fmt.Printf(entry.GetLabel())
	}
}
