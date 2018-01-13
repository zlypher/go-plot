package plot

// Default characters used to draw to various chart components
const (
	xAxisChar     = "-"
	yAxisChar     = "|"
	crossAxisChar = "+"
	barChar       = "+"
)

type Theme struct {
	XAxis     string
	YAxis     string
	CrossAxis string
	Bar       string
}

func DefaultTheme() Theme {
	return Theme{
		XAxis:     xAxisChar,
		YAxis:     yAxisChar,
		CrossAxis: crossAxisChar,
		Bar:       barChar,
	}
}
