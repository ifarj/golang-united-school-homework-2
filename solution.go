package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type sidesNumber int

const (
	SidesTriangle = sidesNumber(3)
	SidesSquare   = sidesNumber(4)
	SidesCircle   = sidesNumber(0)
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum sidesNumber) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (math.Sqrt(3) * sideLen * sideLen) / 4
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * sideLen * sideLen
	}
	return 0
}
