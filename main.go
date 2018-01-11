package main

import (
	"fmt"

	"github.com/zlypher/plot/plot"
)

func main() {
	fmt.Println("Hello World2")

	var entries []plot.Entry
	entries = append(entries, plot.Entry{Label: "Hello", LabelAbbr: "A", XValue: 1, YValue: 1})
	entries = append(entries, plot.Entry{Label: "Hello", LabelAbbr: "B", XValue: 2, YValue: 2})
	entries = append(entries, plot.Entry{Label: "Hello", LabelAbbr: "C", XValue: 3, YValue: 3})
	entries = append(entries, plot.Entry{Label: "Hello", LabelAbbr: "D", XValue: 4, YValue: 0})
	entries = append(entries, plot.Entry{Label: "Hello", LabelAbbr: "E", XValue: 5, YValue: 5})
	entries = append(entries, plot.Entry{Label: "Hello", LabelAbbr: "F", XValue: 6, YValue: 2})

	plot.BarChart(entries)
}
