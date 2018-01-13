package plot

import (
	"fmt"
	"math"
	"strings"

	"github.com/zlypher/go-plot/chart"
)

// Entry represents a single point on the chart.
type Entry struct {
	Label     string
	LabelAbbr string
	XValue    float64
	YValue    float64
}

// Chart holds all required data to render the chart.
type Chart struct {
	Title   string
	Debug   bool
	Spacing Spacing
	Entries []Entry
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
		printDebugInfo(numEntries, width)
	}

	xA := calculateAxis(chart.Entries)
	// TODO: Determine max with of yaxis label
	axisLabelWidth := 5

	if chart.Title != "" {
		printTitle(chart.Title, width+axisLabelWidth+3)
	}
	printChart(chart.Entries, xA, chart.Theme, axisLabelWidth)

	printXAxis(chart.Theme, width, axisLabelWidth)
	printXAxisLabels(chart.Entries, axisLabelWidth)
	fmt.Println()
}

func calculateWidth(sp Spacing, num int) int {
	return sp.Axis + (2 * sp.Margin) + num*sp.Bar + (num-1)*sp.Padding
}

func calculateAxis(entries []Entry) chart.Axis {
	low, high, steps := 0.0, math.SmallestNonzeroFloat64, 0.0

	for _, entry := range entries {
		if entry.YValue < low {
			low = entry.YValue
		} else if entry.YValue > high {
			high = entry.YValue
		}
	}

	steps = 1.0

	return chart.Axis{Low: low, High: high, Steps: steps}
}

func printDebugInfo(numEntries int, width int) {
	fmt.Println("-----")
	fmt.Printf("Number of Entries: %d\n", numEntries)
	fmt.Printf("Width of Chart %d\n", width)
	fmt.Printf("-----\n\n")
}

func printTitle(title string, width int) {
	titleLen := len(title)
	if titleLen >= width {
		fmt.Printf("%s\n\n", title)
		return
	}

	pad := (width - titleLen) / 2

	fmt.Printf("%s%s\n\n",
		strings.Repeat(" ", pad),
		title)
}

func printChart(entries []Entry, axis chart.Axis, theme Theme, axisLabelWidth int) {
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
			if entry.YValue >= val {
				fmt.Printf(theme.Bar)
			} else {
				fmt.Printf(" ")
			}
		}

		// print remaining margin and newline
		fmt.Printf("%s\n", strings.Repeat(" ", 2))
	}
}

func printXAxis(theme Theme, width int, axisLabelWidth int) {
	fmt.Printf("%s%s%s\n",
		strings.Repeat(" ", axisLabelWidth+3),
		theme.CrossAxis,
		strings.Repeat(theme.XAxis, width-1))
}

func printXAxisLabels(entries []Entry, axisLabelWidth int) {
	fmt.Printf("%s%s",
		strings.Repeat(" ", axisLabelWidth+3),
		strings.Repeat(" ", 3)) // axis + margin

	for idx, entry := range entries {
		if idx != 0 {
			fmt.Printf(strings.Repeat(" ", 2)) // pad
		}

		fmt.Printf(entry.LabelAbbr)
	}
}
