// trans3d provides code for representing orientations and performing transformations.
//
// The code is based on the book 3D Math Primer for Graphics and Game Development
// by Fletcher Dunn and Ian Parberry.
package trans3d

import (
	"math"
)

const (
	KPi = math.Pi
	K2Pi = KPi * 2.0
	KPiOver2 = KPi / 2.0
	K1OverPi = 1.0 / KPi
	k1Over2Pi = 1.0 / K2Pi
)

// wrapPi "Wraps" an angle in range -pi--pi by adding the correct multiple of
// 2 pi.
func WrapPi(theta float64) float64 {
	theta += KPi
	theta -= math.Floor(theta * k1Over2Pi) * K2Pi
	theta -= KPi
	return theta
}

// safeAcos is the same as acos(x), but if x is out of range, it is "clamped"
// to the nearest valid value. The value returned is in the range 0..pi, the
// same as the standard Go acos() function.
func SafeAcos(x float64) float64 {

	// Check limit condition

	if (x <= -1.0) {
		return KPi
	}
	if (x >= 1.0) {
		return 0.0
	}

	return math.Acos(x)
}


