package interpolation

type MathFunc func(float64) float64

func Lagrange(xs, ys []float64) (_ MathFunc) {
	if len(xs) != len(ys) || len(xs) == 0 {
		return
	}
	// Копия слайса точек для использовния внутри замыкания
	points := append(xs[:0:0], xs...)

	// Коэффициенты полинома
	coeffs := make([]float64, 0, len(xs))

	for i, c := range points {
		product := 1.0
		for j, n := range points {
			if i != j {
				product *= c - n
			}
		}
		coeffs = append(coeffs, ys[i]/product)
	}

	// Замыкание, ведущее себя как полином Лагранжа
	return func(x float64) float64 {
		res := 0.0
		for i := range coeffs {
			tmp := coeffs[i]
			for j, c := range points {
				if i != j {
					tmp *= x - c
				}
			}
			res += tmp
		}

		return res
	}
}
