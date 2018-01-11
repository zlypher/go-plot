package plot

import (
	"fmt"
	"math"
	"strings"

	"github.com/zlypher/plot/chart"
)

type Entry struct {
	Label     string
	LabelAbbr string
	XValue    float64
	YValue    float64
}

type Chart struct {
	Title   string
	Spacing Spacing
	Entries []Entry
}

type Spacing struct {
	Margin  int
	Padding int
	Bar     int
	Axis    int
}

const (
	xAxisChar = "-"
	yAxisChar = "|"
	spaceChar = " "
	barChar   = "+"
)

func BarChart(chart Chart) {
	numEntries := len(chart.Entries)
	width := calculateWidth(chart.Spacing, numEntries)

	fmt.Printf("Num Entries: %d\n", numEntries)
	fmt.Printf("Width of %d\n", width)

	printBarChart(chart, width)
	fmt.Println()
}

func calculateWidth(sp Spacing, num int) int {
	return sp.Axis + (2 * sp.Margin) + num*sp.Bar + (num-1)*sp.Padding
}

func calculateAxis(entries []Entry) chart.Axis {
	low, high, steps := math.MaxFloat64, math.SmallestNonzeroFloat64, 0.0

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

/**
margin      bar        pad        bar        margin
xxxx    +    x    +    xx    +    x     +    xxxx


          +
       +  +
    +  +  +
---------------

*/

func printBarChart(chrt Chart, width int) {
	printTitle()

	xA := calculateAxis(chrt.Entries)
	printChart(chrt.Entries, xA)

	printXAxis(width)
	printXAxisLabels(chrt.Entries)
}

func printTitle() {
	fmt.Println("            ++ Lorem Ipsum ++             ")
	fmt.Println()
}

func printChart(entries []Entry, xAxis chart.Axis) {
	fmt.Println(yAxisChar) // empty line
	for i := xAxis.High; i >= xAxis.Low; i -= xAxis.Steps {

		fmt.Printf(yAxisChar)
		fmt.Printf(strings.Repeat(spaceChar, 2)) // margin

		for idx, entry := range entries {
			if idx != 0 {
				fmt.Printf(strings.Repeat(spaceChar, 2)) // pad
			}

			if entry.YValue >= i {
				fmt.Printf(barChar)
			} else {
				fmt.Printf(spaceChar)
			}
		}

		fmt.Printf(strings.Repeat(spaceChar, 2)) // margin
		fmt.Println()
	}
}

func printXAxis(width int) {
	fmt.Println(strings.Repeat(xAxisChar, width))
}

func printXAxisLabels(entries []Entry) {
	fmt.Printf(strings.Repeat(spaceChar, 1+2)) // axis + margin

	for idx, entry := range entries {
		if idx != 0 {
			fmt.Printf(strings.Repeat(spaceChar, 2)) // pad
		}

		fmt.Printf(entry.LabelAbbr)
	}

	fmt.Println()
}

func printLegend() {
	fmt.Println("TODO")
}
