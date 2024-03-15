package main

import (
	"fmt"
	"math"
)

type MathFunc func(x, y float64) float64

type Pair struct {
	x, y float64
}

type Table []Pair

func (t Table) String() string {
	str := ""
	for i := range t {
		str += fmt.Sprintf("%3d % -5.5v % -5.5v\n", i, t[i].x, t[i].y)
	}

	return str
}

func Euler(x0, y0, xn, h float64, f MathFunc) Table {
	if x0 > xn {
		return nil
	}

	answer := make(Table, int((xn-x0)/h)+1)

	x := x0
	y := y0
	for i := 0; ; {
		answer[i].x = x
		answer[i].y = y

		i++

		if i >= len(answer) {
			break
		}

		y += h * f(x, y)
		x += h
	}

	return answer
}

func EulerImplicit(x0, y0, xn, h float64, f MathFunc, r func(x float64) float64) Table {
	if x0 > xn {
		return nil
	}

	answer := make(Table, int((xn-x0)/h)+1)

	x := x0
	y := y0
	for i := 0; ; {
		answer[i].x = x
		answer[i].y = y

		i++

		if i >= len(answer) {
			break
		}

		y += h * f(x, r(x))
		x += h
	}

	return answer
}

func EulerImproved(x0, y0, xn, h float64, f MathFunc) Table {
	if x0 > xn {
		return nil
	}

	answer := make(Table, int((xn-x0)/h)+1)

	x := x0
	y := y0
	for i := 0; ; {
		answer[i].x = x
		answer[i].y = y

		i++

		if i >= len(answer) {
			break
		}

		halfX := x + h/2
		halfY := y + h/2*f(x, y)

		y += h * f(halfX, halfY)
		x += h
	}

	return answer
}

func EulerCauchy(x0, y0, xn, h float64, f MathFunc) Table {
	if x0 > xn {
		return nil
	}

	answer := make(Table, int((xn-x0)/h)+1)

	x := x0
	y := y0
	for i := 0; ; {
		answer[i].x = x
		answer[i].y = y

		i++

		if i >= len(answer) {
			break
		}

		yi := y + h*f(x, y)
		for range 4 {
			yi = y + 0.5*h*(f(x, y)+f(x+h, yi))
		}
		y = yi
		x += h
	}
	return answer
}

func printEpsilon(answer Table, r func(x float64) float64) {
	fmt.Println("epsilon:")
	for i := range answer {
		x := answer[i].x
		y := answer[i].y
		fmt.Printf("%3d % -5.5e\n", i, math.Abs(r(x)-y))
	}
}

func main() {
	// Вариант 14.
	x0 := 1.0
	y0 := 1.0
	xn := 2.0
	h := 0.1

	f := func(x, y float64) float64 {
		return -0.5*y/x + x*x
	}

	r := func(x float64) float64 {
		return 2.0/7.0*x*x*x + 5.0/(7.0*math.Sqrt(x))
	}

	// Метод Эйлера
	answer := Euler(x0, y0, xn, h, f)

	fmt.Println("Метод Эйлера:")
	fmt.Println(answer)

	printEpsilon(answer, r)

	// Неявный метод Эйлера
	answer = EulerImplicit(x0, y0, xn, h, f, r)
	fmt.Println()
	fmt.Println("Неявный метод Эйлера")
	fmt.Println(answer)

	printEpsilon(answer, r)

	// Метод Эйлера-Коши с итерационной обработкой
	answer = EulerCauchy(x0, y0, xn, h, f)

	fmt.Println()
	fmt.Println("Метод Эйлера-Коши с итерационной обработкой:")
	fmt.Println(answer)

	printEpsilon(answer, r)

	// Улучшенный метод Эйлера
	answer = EulerImproved(x0, y0, xn, h, f)

	fmt.Println()
	fmt.Println("Улучшенный метод Эйлера:")
	fmt.Println(answer)

	printEpsilon(answer, r)
}
