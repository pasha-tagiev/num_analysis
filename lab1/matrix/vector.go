package matrix

import "math"

type Vector []float64

func (v Vector) Copy() Vector {
	copy := append(make(Vector, 0, len(v)), v...)
	return copy
}

func (v Vector) Len() int {
	return len(v)
}

func (v Vector) Sum(vec Vector) {
	for i := range v {
		v[i] += vec[i]
	}
}

func (v Vector) Sub(vec Vector) {
	for i := range v {
		v[i] -= vec[i]
	}
}

func (v Vector) Mul(vec Vector) float64 {
	sum := .0
	for i := range v {
		sum += v[i] * vec[i]
	}

	return sum
}

func (v Vector) Norm1() float64 {
	norm := .0
	for _, e := range v {
		norm += math.Abs(e)
	}

	return norm
}

func (v Vector) NormC() float64 {
	norm := .0
	for _, e := range v {
		curr := math.Abs(e)
		if norm < curr {
			norm = curr
		}
	}

	return norm
}
