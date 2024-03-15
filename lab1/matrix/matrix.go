package matrix

import (
	"fmt"
	"math"
	"strings"
)

type Matrix []Vector

func New(m, n int) Matrix {
	mx := make(Matrix, m)
	for i := range mx {
		mx[i] = make(Vector, n)
	}
	return mx
}

func FromSlice(m, n int, init []float64) Matrix {
	if m*n != len(init) {
		return nil
	}
	mx := make(Matrix, 0, m)
	for len(init) != 0 {
		init, mx = init[n:], append(mx, init[:n])
	}
	return mx
}

func (mx Matrix) Diagonal() []float64 {
	m := min(mx.M(), mx.N())
	res := make([]float64, m)
	for i := range m {
		res[i] = mx[i][i]
	}
	return res
}

func E(m int) Matrix {
	mx := New(m, m)
	for i := range mx.M() {
		mx[i][i] = 1.0
	}
	return mx
}

func (mx Matrix) Copy() Matrix {
	copy := make(Matrix, mx.M())
	for i := range copy {
		copy[i] = mx[i].Copy()
	}
	return copy
}

func (mx Matrix) NormC() float64 {
	norm := .0
	for _, row := range mx {
		sum := .0
		for _, elm := range row {
			sum += math.Abs(elm)
		}
		if norm < sum {
			norm = sum
		}
	}

	return norm
}

func (mx Matrix) Sub(rhv Matrix) Matrix {
	res := New(mx.M(), mx.N())
	for m, row := range res {
		for n := range row {
			row[n] = mx[m][n] - rhv[m][n]
		}
	}
	return res
}

func (mx Matrix) Mul(rhv Matrix) Matrix {
	res := New(mx.M(), rhv.N())
	for m := range res.M() {
		for n := range res.N() {
			for i := range mx.N() {
				res[m][n] += mx[m][i] * rhv[i][n]
			}
		}
	}
	return res
}

func (mx Matrix) MulByNum(num float64) Matrix {
	res := New(mx.M(), mx.N())
	for m, row := range res {
		for n := range row {
			row[n] = num * mx[m][n]
		}
	}
	return res
}

func (mx Matrix) MulByVec(vec Vector) Vector {
	if mx.N() != len(vec) {
		return nil
	}

	res := make(Vector, len(vec))

	i := 0
	for m, row := range mx {
		for n := range row {
			res[i] += mx[m][n] * vec[n]
		}
		i++
	}

	return res
}

func (mx Matrix) Transpose() Matrix {
	res := New(mx.N(), mx.M())
	for m := range res.M() {
		for n := range res.N() {
			res[m][n] = mx[n][m]
		}
	}
	return res
}

func (mx Matrix) M() int {
	return len(mx)
}

func (mx Matrix) N() int {
	if mx.M() == 0 {
		return 0
	}
	return len(mx[0])
}

func (mx Matrix) IsSquare() bool {
	return mx.M() == mx.N()
}

func (mx Matrix) String() string {
	var builder strings.Builder

	for _, row := range mx {
		for _, elm := range row {
			builder.WriteString(fmt.Sprintf("% -15.3f", elm))
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}
