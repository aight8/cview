package main

import (
	"fmt"

	"github.com/aight8/cview"
	"github.com/gdamore/tcell/v2"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content cview.Primitive) {
	textView := cview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/aight8/cview"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
