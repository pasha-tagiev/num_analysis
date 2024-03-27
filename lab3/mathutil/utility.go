package mathutil

import (
	"math"
)

// Машинный ноль для float32 и float64(double).
var epsilon32 = math.Nextafter32(1, 2) - 1
var epsilon64 = math.Nextafter(1, 2) - 1

// Констрейнт для чисел с плавающей точкой для использования в дженериках.
type Float interface {
	float32 | float64
}

func Epsilon[F Float]() (eps F) {
	switch any(eps).(type) {
	case float32:
		eps = F(epsilon32)
	case float64:
		eps = F(epsilon64)
	}
	return
}

func Equal[F Float](a, b F) bool {
	return F(math.Abs(float64(a-b))) <= Epsilon[F]()
}

// Компаратор для функций стандартной библиотеки.
func CmpFloats[F Float](a, t F) int {
	if Equal(a, t) {
		return 0
	}

	if a < t {
		return -1
	}

	return 1
}
