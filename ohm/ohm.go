package ohm

func Ohm(v, r, i float32) float32 {

	if v > 0 && r > 0 && i > 0 {
		return 0
	} else if v > 0 && r > 0 {
		return v / r
	} else if v > 0 && i > 0 {
		return v / i
	} else if i > 0 && r > 0 {
		return i * r
	} else {
		return 0
	}
}