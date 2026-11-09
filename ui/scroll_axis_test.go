package ui

import (
	"testing"
	"time"

	"github.com/egoist/mygo/internal/platform"
)

// A page scrolling down over rows that scroll sideways, as a media
// library's shelves: a touchpad's gesture goes along the axis it mostly
// goes, as AppKit's do, so the page scrolls down under a slightly diagonal
// swipe rather than the row taking it.
func TestTouchpadGestureKeepsItsAxis(t *testing.T) {
	var page, row ScrollState
	shelf := func(c *context) {
		coreScroll(c).Fill().TrackScroll(&page).Children(func() {
			coreScrollHorizontal(c).Height(100).Shrink(0).TrackScroll(&row).Children(func() {
				coreBox(c).Width(2000).Height(100).Shrink(0)
			})
			coreBox(c).Height(2000).Shrink(0)
		})
	}
	tt := coreNewTester(shelf, 400, 300)
	tt.Move(200, 50)
	// Down, drifting a little sideways, over the row.
	for range 10 {
		tt.Scroll(200, 50, 1.5, 12)
		tt.Advance(16 * time.Millisecond)
	}
	if page.Y != 120 || row.X != 0 {
		t.Errorf("a swipe down over the row: page at %v, row at %v; want 120 and 0", page.Y, row.X)
	}
	// After a pause, a swipe sideways, drifting a little down, scrolls the
	// row alone.
	tt.Advance(time.Second)
	page.Y = 0
	tt.Frame()
	for range 10 {
		tt.Scroll(200, 50, 12, 1.5)
		tt.Advance(16 * time.Millisecond)
	}
	if row.X != 120 || page.Y != 0 {
		t.Errorf("a swipe sideways over the row: row at %v, page at %v; want 120 and 0", row.X, page.Y)
	}
	// The momentum after the fingers keeps the gesture's axis.
	tt.send(platform.SurfaceEvent{Kind: platform.PointerScroll, X: 200, Y: 50, DX: 3, DY: 30, Precise: true, Momentum: true})
	if page.Y != 0 || row.X != 123 {
		t.Errorf("momentum of the sideways swipe: row at %v, page at %v; want 123 and 0", row.X, page.Y)
	}
}
