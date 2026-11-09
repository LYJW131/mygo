package ui

import "testing"

// A video player's picture: a click pauses, a double click goes full
// screen, and a press that moves drags the window, clicking nothing.
func TestDragWindowOnMove(t *testing.T) {
	clicks := 0
	tt := coreNewTester(func(c *context) {
		b := coreBox(c).Fill().DragWindowOnMove()
		if b.Clicked() {
			clicks++
		}
	}, 300, 200)
	tt.Move(100, 100)
	tt.Press(100, 100)
	tt.Release(100, 100)
	tt.Frame()
	if clicks != 1 || tt.h.windowDrags != 0 {
		t.Fatalf("a click: %d clicks, %d window drags; want 1 and 0", clicks, tt.h.windowDrags)
	}
	tt.Press(100, 100)
	tt.Move(102, 101)
	if tt.h.windowDrags != 0 {
		t.Fatalf("a press that moved 2 DIPs drags the window")
	}
	tt.Move(110, 104)
	tt.Release(110, 104)
	tt.Frame()
	if clicks != 1 || tt.h.windowDrags != 1 {
		t.Errorf("a press that moved: %d clicks, %d window drags; want 1 and 1", clicks, tt.h.windowDrags)
	}
}
