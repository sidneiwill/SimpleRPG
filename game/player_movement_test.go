package game

import (
	"math"
	"testing"
)

func TestMovementDelta(t *testing.T) {
	const speed = 2.0
	diagonal := speed / math.Sqrt2

	tests := []struct {
		name                  string
		up, down, left, right bool
		wantX, wantY          float64
	}{
		{
			name:  "idle",
			wantX: 0,
			wantY: 0,
		},
		{
			name:  "right",
			right: true,
			wantX: speed,
			wantY: 0,
		},
		{
			name:  "up",
			up:    true,
			wantX: 0,
			wantY: -speed,
		},
		{
			name:  "up right",
			up:    true,
			right: true,
			wantX: diagonal,
			wantY: -diagonal,
		},
		{
			name: "down left",
			down: true, left: true,
			wantX: -diagonal, wantY: diagonal,
		},
		{
			name:  "horizontal directions cancel",
			left:  true,
			right: true,
			wantX: 0,
			wantY: 0,
		},
		{
			name:  "vertical direction cancel",
			up:    true,
			down:  true,
			wantX: 0,
			wantY: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := movementDelta(
				tt.up,
				tt.down,
				tt.left,
				tt.right,
				speed,
			)

			const tolerance = 1e-9

			if math.Abs(gotX-tt.wantX) > tolerance {
				t.Errorf("x = %v, want %v", gotX, tt.wantX)
			}

			if math.Abs(gotY-tt.wantY) > tolerance {
				t.Errorf("y = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
