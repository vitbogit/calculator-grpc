package math

// Наибольший общий делитель (greatest common divisor)
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Наименьшее общее кратное (least common multiple)
func LCM(a, b int64) int64 {
	return Modulus(a*b) / GCD(a, b)
}

func LCMWithCoeffs(a, b int64) (int64, int64, int64) {
	LCM := LCM(a, b)
	return LCM, LCM / a, LCM / b
}
