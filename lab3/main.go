package main

import (
	"fmt"
	"math"
	"num_analysis/lab3/interpolation"
	"num_analysis/lab3/mathutil"
	"slices"
)

type MathFunc func(x float64) float64
type Equation func(x, y float64, yf MathFunc) float64

func Euler(x0, xn, y0, h float64, eq Equation) (_, _ []float64) {
	size := int((xn-x0)/h) + 1
	xs := make([]float64, 0, size)
	ys := make([]float64, 0, size)

	// Функция y в которой производится выбор значения по описанным в теории правилам.
	yf := func(x float64) float64 {
		// CmpFloats - компаратор для бинарного поиска чисел с плавающей точкой.
		i, exists := slices.BinarySearchFunc(xs, x, mathutil.CmpFloats)
		if exists || i == 0 {
			return ys[i]
		}

		// Если x лежит правее x0 и он еще не встречался,
		// производится интеполяция полиномом Лагранжа.
		// Реализация в interpolation/lagrange.go.
		L2 := interpolation.Lagrange(xs[i-1:i+2], ys[i-1:i+2])
		return L2(x)
	}

	x := x0
	y := y0
	for range size {
		xs = append(xs, x)
		ys = append(ys, y)

		xhalf := x + h/2
		yhalf := y + h/2*eq(x, y, yf)

		x += h
		y += h * eq(xhalf, yhalf, yf)
	}

	return xs, ys
}

func main() {
	// Задание 4.3. Вариант 14.
	A1 := 1.3
	A2 := 2.0
	A3 := 0.23

	// Уравнение 14 варианта.
	eq := func(x, y float64, yf MathFunc) float64 {
		return A2 - A1*y - math.Pow(yf(x-A3), 2)
	}

	h := 0.1
	x0 := 0.0
	xn := 1.0
	y0 := 0.3

	xs, ys := Euler(x0, xn, y0, h, eq)

	// Ответы на скрине.
	for i := range xs {
		fmt.Printf("%-4.2v %-10.6v\n", xs[i], ys[i])
	}
}
