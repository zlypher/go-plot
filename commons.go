package plot

import (
	"fmt"
	"io"
	"strings"
)

func formatDebugInfo(numEntries int, width int) string {
	return fmt.Sprint(
		fmt.Sprintf("-----\n"),
		fmt.Sprintf("Number of Entries: %d\n", numEntries),
		fmt.Sprintf("Width of Chart %d\n", width),
		fmt.Sprintf("-----\n\n"),
	)
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
