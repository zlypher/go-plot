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
}

// Spacing defines sizes of various spacing elements (margin, padding, ...).
type Spacing struct {
	Margin  int
	Padding int
	Bar     int
	Axis    int
}

// Characters used to draw to various chart components
const (
	xAxisChar     = "-"
	yAxisChar     = "|"
	crossAxisChar = "+"
	spaceChar     = " "
	barChar       = "+"
)

// BarChart draws the bar chart to the cmd.
func BarChart(chart Chart) {
	numEntries := len(chart.Entries)
	width := calculateWidth(chart.Spacing, numEntries)

	if chart.Debug {
		printDebugInfo(numEntries, width)
	}

	printBarChart(chart, width)
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

func printBarChart(chrt Chart, width int) {
	xA := calculateAxis(chrt.Entries)
	// TODO: Determine max with of yaxis label
	axisLabelWidth := 5

	if chrt.Title != "" {
		printTitle(chrt.Title, width+axisLabelWidth+3)
	}
	printChart(chrt.Entries, xA, axisLabelWidth)

	printXAxis(width, axisLabelWidth)
	printXAxisLabels(chrt.Entries, axisLabelWidth)
}

func printTitle(title string, width int) {
	titleLen := len(title)
	if titleLen >= width {
		fmt.Printf("%s\n\n", title)
		return
	}

	pad := (width - titleLen) / 2

	fmt.Printf("%s%s%s\n\n",
		strings.Repeat(spaceChar, pad),
		title,
		strings.Repeat(spaceChar, pad))
}

func printChart(entries []Entry, axis chart.Axis, axisLabelWidth int) {
	// Start the chart with a line with only the y axis drawn
	fmt.Printf("%s%s\n", strings.Repeat(spaceChar, axisLabelWidth+3), yAxisChar)

	labelFmt := fmt.Sprintf("%%%d.0f - ", axisLabelWidth)
	for val := axis.High; val >= axis.Low; val -= axis.Steps {
		// Print current y axis value
		fmt.Printf(labelFmt, val)

		// Print the y axis and the margin until the first bar
		fmt.Printf("%s%s", yAxisChar, strings.Repeat(spaceChar, 2))

		// Print the bars with padding between each bar
		for idx, entry := range entries {
			// If it is not the first element, draw the padding
			if idx != 0 {
				fmt.Printf(strings.Repeat(spaceChar, 2))
			}

			// If the bar reaches up to the current value, draw the bar.
			// If not, draw a spacing.
			if entry.YValue >= val {
				fmt.Printf(barChar)
			} else {
				fmt.Printf(spaceChar)
			}
		}

		// print remaining margin and newline
		fmt.Printf("%s\n", strings.Repeat(spaceChar, 2))
	}
}

func printXAxis(width int, axisLabelWidth int) {
	fmt.Printf("%s%s%s\n",
		strings.Repeat(spaceChar, axisLabelWidth+3),
		crossAxisChar,
		strings.Repeat(xAxisChar, width-1))
}

func printXAxisLabels(entries []Entry, axisLabelWidth int) {
	fmt.Printf("%s%s",
		strings.Repeat(spaceChar, axisLabelWidth+3),
		strings.Repeat(spaceChar, 1+2)) // axis + margin

	for idx, entry := range entries {
		if idx != 0 {
			fmt.Printf(strings.Repeat(spaceChar, 2)) // pad
		}

		fmt.Printf(entry.LabelAbbr)
	}
}
