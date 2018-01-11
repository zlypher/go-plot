package plot

import (
	"fmt"
	"strings"
)

type Entry struct {
	Label     string
	LabelAbbr string
	XValue    float32
	YValue    float32
}

type Chart struct {
	Title   string
	Spacing Spacing
	Entries []Entry
}

type Spacing struct {
	Margin  int
	Padding int
}

type axis struct {
	low   float32
	high  float32
	steps float32
}

const (
	xAxisChar = "-"
	yAxisChar = "|"
	spaceChar = " "
	barChar   = "+"
)

func BarChart(entries []Entry) {
	fmt.Println("Bar Chart")
	fmt.Println()

	axisWidth, margin, pad, barWidth := 1, 2, 2, 1

	numEntries := len(entries)
	width := axisWidth + (2 * margin) + numEntries*barWidth + (numEntries-1)*pad

	fmt.Printf("Num Entries: %d\n", numEntries)
	fmt.Printf("Width of %d\n", width)

	printBarChart(entries, width)
	fmt.Println()
}

/**
margin      bar        pad        bar        margin
xxxx    +    x    +    xx    +    x     +    xxxx


          +
       +  +
    +  +  +
---------------

*/

func printBarChart(entries []Entry, width int) {
	printTitle()

	xA := axis{low: 1, high: 5, steps: 1}
	fmt.Println(xA)

	printChart(entries, xA)

	printXAxis(width)
	printXAxisLabels(entries)
}

func printTitle() {
	fmt.Println("            ++ Lorem Ipsum ++             ")
	fmt.Println()
}

func printChart(entries []Entry, xAxis axis) {
	fmt.Println(yAxisChar) // empty line
	for i := xAxis.high; i >= xAxis.low; i -= xAxis.steps {

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
