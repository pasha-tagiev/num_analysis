package matrix

import (
	"errors"
	"math"
)

var (
	ErrZeroOnDiagonal       = errors.New("zero on diagonal")
	ErrNormIsGreaterThanOne = errors.New("norm is greater than one")
	ErrIncorrectArgument    = errors.New("incorrect argument")
)

func epsilon() float64 {
	return math.Nextafter(1, 2) - 1
}

func transform(mx Matrix, b Vector) error {
	for m := range mx {
		c := mx[m][m]

		if math.Abs(c) < epsilon() {
			return ErrZeroOnDiagonal
		}

		b[m] /= c
		for n := 0; n < mx.N(); n++ {
			mx[m][n] /= -c
		}
		mx[m][m] = 0
	}

	return nil
}

func prepare(mx Matrix, b Vector) (norm float64, err error) {
	if !mx.IsSquare() || mx.M() != b.Len() {
		err = ErrIncorrectArgument
		return
	}

	if err = transform(mx, b); err != nil {
		return
	}

	norm = mx.NormC()
	if norm >= 1 {
		err = ErrNormIsGreaterThanOne
		return
	}

	return
}

func currentEps(coef float64, curr, prev Vector) (float64, Vector) {
	prev.Sub(curr)
	return coef * prev.NormC(), curr
}

func FixedPoint(mx Matrix, b Vector, eps float64, iterLimit int) (x Vector, iterCount int, err error) {
	aNorm, err := prepare(mx, b)
	if err != nil {
		return
	}

	coef := aNorm / (1 - aNorm)

	x = b.Copy()
	for ; iterCount < iterLimit; iterCount++ {
		n := mx.MulByVec(x)
		n.Sum(b)

		cureps := .0
		if cureps, x = currentEps(coef, n, x); eps >= cureps {
			break
		}
	}

	return
}

// Норма верхней треугольной матрицы
func upperMatrixNorm(mx Matrix) float64 {
	norm := .0
	for i, row := range mx {
		curr := row[i:].Norm1()
		if norm < curr {
			norm = curr
		}
	}

	return norm
}

func Seidel(mx Matrix, b Vector, eps float64, iterLimit int) (x Vector, iterCount int, err error) {
	aNorm, err := prepare(mx, b)
	if err != nil {
		return
	}

	coef := upperMatrixNorm(mx) / (1 - aNorm)

	x = b.Copy()
	for ; iterCount < iterLimit; iterCount++ {
		n := x.Copy()
		for m, row := range mx {
			n[m] = row.Mul(n) + b[m] // row.Mul(n) - скалярное произведение строки матрицы на вектор n
		}

		cureps := .0
		if cureps, x = currentEps(coef, n, x); eps >= cureps {
			break
		}
	}

	return
}

