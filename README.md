# go-plot

go-plot is a utility library to display simplistic charts in a command line.

[![Build Status](https://travis-ci.org/zlypher/go-plot.svg?branch=master)](https://travis-ci.org/zlypher/go-plot)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**Table of Contents**
* [Getting Started](#getting-started)
* [Example](#example)
* [Contributing](#contributing)
* [License](#license)

## Getting Started

To get started, simply `go get` the library and import it into your file.

```bash
go get github.com/zlypher/go-plot
```

```go
import (
    "github.com/zlypher/go-plot"
)
```

## Example

**main.go**
```go
package main

import (
    "github.com/zlypher/go-plot"
)

func main() {
    var entries []plot.Entry
    entries = append(entries, plot.Entry{Label: "Lorem Ipsum", LabelAbbr: "A", XValue: 1, YValue: 1})
    entries = append(entries, plot.Entry{Label: "Lorem Ipsum", LabelAbbr: "B", XValue: 2, YValue: 2})
    entries = append(entries, plot.Entry{Label: "Lorem Ipsum", LabelAbbr: "C", XValue: 3, YValue: 3})

    spacing := plot.Spacing{Margin: 2, Padding: 2, Bar: 1, Axis: 1}

    chart := plot.Chart{
        Title:   "",
        Debug:   false,
        Spacing: spacing,
        Entries: entries,
    }

    plot.BarChart(chart)
}
```

**Example Output**
```bash
|
|        +
|     +  +
|  +  +  +
+-----------
   A  B  C
```

## Contributing

This is a small side project, so there currently are not strict standards for contributing. If you find any issues or have suggestions, feel free to [create an issue](https://github.com/zlypher/bool/issues) or a [pull request](https://github.com/zlypher/bool/pulls)

## License

MIT License (see [LICENSE.md](LICENSE.md))
