package main

import (
	"testing"

	"github.com/Zyko0/go-sdl3/sdl"
)

// TestShrinkRectKeepsCenter checks that shrinking a rect scales it about its
// own center, so a timer shrinks in place instead of drifting towards a corner.
func TestShrinkRectKeepsCenter(t *testing.T) {
	r := sdl.FRect{X: 530, Y: 10, W: 1380, H: 300}
	centerX := r.X + r.W/2
	centerY := r.Y + r.H/2

	shrinkRect(&r, 0.5)

	if r.W != 690 || r.H != 150 {
		t.Errorf("expected 690x150, got %vx%v", r.W, r.H)
	}
	if got := r.X + r.W/2; got != centerX {
		t.Errorf("center X moved: expected %v, got %v", centerX, got)
	}
	if got := r.Y + r.H/2; got != centerY {
		t.Errorf("center Y moved: expected %v, got %v", centerY, got)
	}
}

// TestShrinkRectNoOps checks the cases that must leave the rect untouched, so
// that the default scale of 1.0 renders byte-identically to the old layout.
func TestShrinkRectNoOps(t *testing.T) {
	tests := []struct {
		name  string
		rect  sdl.FRect
		scale float32
	}{
		{"scale of 1.0", sdl.FRect{X: 25, Y: 290, W: 1870, H: 440}, 1.0},
		{"scale above 1.0", sdl.FRect{X: 25, Y: 290, W: 1870, H: 440}, 1.5},
		{"zero scale", sdl.FRect{X: 25, Y: 290, W: 1870, H: 440}, 0},
		{"negative scale", sdl.FRect{X: 25, Y: 290, W: 1870, H: 440}, -0.5},
		{"empty rect", sdl.FRect{X: 25, Y: 290, W: 0, H: 0}, 0.5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.rect
			shrinkRect(&got, tc.scale)
			if got != tc.rect {
				t.Errorf("rect changed: expected %+v, got %+v", tc.rect, got)
			}
		})
	}
}

// TestTextClockScaleClamps checks that a hand-edited clock.ini cannot push the
// scale outside the range the web UI enforces, and that an unset option renders
// the historical layout rather than being clamped up to the minimum.
func TestTextClockScaleClamps(t *testing.T) {
	tests := []struct {
		configured float64
		want       float32
	}{
		{1.0, 1.0},
		{0.5, 0.5},
		{0.75, 0.75},
		{2.0, maxTextClockScale},
		{0.1, minTextClockScale},
		{0, maxTextClockScale},  // unset: must not shrink anything
		{-1, maxTextClockScale}, // nonsense: must not shrink anything
	}

	original := options.TextClockScale
	defer func() { options.TextClockScale = original }()

	for _, tc := range tests {
		options.TextClockScale = tc.configured
		if got := textClockScale(); got != tc.want {
			t.Errorf("TextClockScale=%v: expected %v, got %v", tc.configured, tc.want, got)
		}
	}
}
