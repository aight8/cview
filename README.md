# cview - Terminal-based user interface toolkit

This package is a fork of [tview](https://github.com/rivo/tview).
See [FORK.md](https://github.com/aight8/cview/src/branch/master/FORK.md) for more information.

## Features

Available widgets:

- __Input forms__ (including __input/password fields__, __drop-down selections__, __checkboxes__, and __buttons__)
- Navigable multi-color __text views__
- Selectable __lists__ with __context menus__
- Modal __dialogs__
- Horizontal and vertical __progress bars__
- __Grid__, __Flexbox__ and __tabbed panel layouts__
- Sophisticated navigable __table views__
- Flexible __tree views__
- Draggable and resizable __windows__
- An __application__ wrapper

Widgets may be customized and extended to suit any application.

[Mouse support](https://docs.rocketnine.space/github.com/aight8/cview#hdr-Mouse_Support) is available.

## Applications

A list of applications powered by cview is available via [pkg.go.dev](https://pkg.go.dev/github.com/aight8/cview?tab=importedby).

## Installation

```bash
go get github.com/aight8/cview
```

## Hello World

This basic example creates a TextView titled "Hello, World!" and displays it in your terminal:

```go
package main

import (
	"github.com/aight8/cview"
)

func main() {
	app := cview.NewApplication()

	tv := cview.NewTextView()
	tv.SetBorder(true)
	tv.SetTitle("Hello, world!")
	tv.SetText("Lorem ipsum dolor sit amet")
	
	app.SetRoot(tv, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
```

Examples are available via [godoc](https://docs.rocketnine.space/github.com/aight8/cview#pkg-examples)
and in the [demos](https://github.com/aight8/cview/src/branch/master/demos) directory.

For a presentation highlighting the features of this package, compile and run
the program in the [demos/presentation](https://github.com/aight8/cview/src/branch/master/demos/presentation) directory.

## Documentation

Package documentation is available via [godoc](https://docs.rocketnine.space/github.com/aight8/cview).

An [introduction tutorial](https://rocketnine.space/post/tview-and-you/) is also available.

## Dependencies

This package is based on [github.com/gdamore/tcell](https://github.com/gdamore/tcell)
(and its dependencies) and [github.com/rivo/uniseg](https://github.com/rivo/uniseg).

## Support

[CONTRIBUTING.md](https://github.com/aight8/cview/src/branch/master/CONTRIBUTING.md) describes how to share
issues, suggestions and patches (pull requests).
