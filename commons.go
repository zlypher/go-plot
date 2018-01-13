package plot

import (
	"fmt"
	"io"
	"strings"
)

func printDebugInfo(w io.Writer, numEntries int, width int) {
	fmt.Fprintln(w, "-----")
	fmt.Fprintf(w, "Number of Entries: %d\n", numEntries)
	fmt.Fprintf(w, "Width of Chart %d\n", width)
	fmt.Fprintf(w, "-----\n\n")
}

func formatTitle(title string, width int) string {
	titleLen := len(title)
	if titleLen >= width {
		return fmt.Sprintf("%s\n\n", title)
	}

	pad := (width - titleLen) / 2

	return fmt.Sprintf("%s%s\n\n",
		strings.Repeat(" ", pad),
		title)
}

func calculateWidth(sp Spacing, num int) int {
	return sp.Axis + (2 * sp.Margin) + num*sp.Bar + (num-1)*sp.Padding
}

// print simply writes the given output to the given writer
func print(w io.Writer, output string) {
	fmt.Fprint(w, output)
}
