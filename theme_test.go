package plot_test

import (
	"testing"

	"github.com/zlypher/go-plot"
)

func TestDefaultTheme(t *testing.T) {
	theme := plot.DefaultTheme()
	expected := plot.Theme{
		XAxis:     "-",
		YAxis:     "|",
		CrossAxis: "+",
		Bar:       "+",
	}

	if theme.XAxis != expected.XAxis {
		t.Errorf("Symbol for XAxis in default theme was incorrect, got: %s, want: %s.", theme.XAxis, expected.XAxis)
	}

	if theme.YAxis != expected.YAxis {
		t.Errorf("Symbol for YAxis in default theme was incorrect, got: %s, want: %s.", theme.YAxis, expected.YAxis)
	}

	if theme.CrossAxis != expected.CrossAxis {
		t.Errorf("Symbol for CrossAxis in default theme was incorrect, got: %s, want: %s.", theme.CrossAxis, expected.CrossAxis)
	}

	if theme.Bar != expected.Bar {
		t.Errorf("Symbol for Bar in default theme was incorrect, got: %s, want: %s.", theme.Bar, expected.Bar)
	}
}
