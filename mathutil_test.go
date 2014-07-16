package trans3d

import (
	"testing"
	"math"
)

func TestWrapPi(t *testing.T) {
	const in = -3 * KPi
	const out = -KPi 
	if x := WrapPi(in); math.Abs(x - out) > 1e-9 {
		t.Errorf("WrapPi(%v) = %v, want %v", in, x, out)
	}
}

func TestSafeAcos(t *testing.T) {
	in := []float64{-2.0, 2.0, 0.5}
	out := []float64{KPi, 0.0, math.Acos(0.5)}
	for i, _ := range in {
		x := SafeAcos(in[i])
		delta := math.Abs(x - out[i])
		if delta > 1e-9 {
			t.Errorf("SafeAcos(%v) = %v, want %v", in[i], x, out[i])
		}
	} 
}
