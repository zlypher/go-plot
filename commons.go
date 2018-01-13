package plot

import (
	"fmt"
	"strings"
)

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
