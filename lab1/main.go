package main

import (
	"fmt"
	"math"
	"num_analysis/lab1/matrix"
)

func PrintAnswer(answer matrix.Vector, correct matrix.Vector) {
	for i, e := range answer {
		fmt.Printf("x%d = % -20v (правильный ответ: % v)\n", i+1, e, correct[i])
	}
}

func main() {
	// Вариант 14. Метод простых итераций и метод Зейдаля

	mx := matrix.FromSlice(4, 4, []float64{
		-22, -2, -6, 6,
		3, -17, -3, 7,
		2, 6, -17, 5,
		-1, -8, 8, 23,
	})

	b := matrix.Vector{96, -26, 35, -234}

	eps := 0.01
	lim := 100

	correct := []float64{-5, -2, -6, -9}

	answer, iterCount, _ := matrix.FixedPoint(mx.Copy(), b.Copy(), eps, lim)
	fmt.Printf("epsilon: %v\n\n", eps)
	fmt.Printf("Метод простых итераций (Кол-во итераций: %d):\n", iterCount)
	PrintAnswer(answer, correct)

	answer, iterCount, _ = matrix.Seidel(mx, b, eps, lim)
	fmt.Println()
	fmt.Printf("Метод Зейделя (Кол-во итераций: %d):\n", iterCount)
	PrintAnswer(answer, correct)

	fmt.Println()

	// Метод вращений

	mx = matrix.FromSlice(3, 3, []float64{
		-7, -5, -9,
		-5, 5, 2,
		-9, 2, 9,
	})

	eps = 0.01

	res, vecs := matrix.Jacobi(mx, eps, lim)

	fmt.Println("Метод вращений")

	fmt.Println("Искомые собственные значения:")
	fmt.Println(res)

	fmt.Println()

	fmt.Println("Собственные векторы:")
	fmt.Println(vecs)

	// Степенной метод. Вычисление спектрального радиуса

	y := matrix.Vector{1, 1, 1}

	pMx := 0.0
	for i := range mx.M() {
		cur := matrix.PowerIteration(mx, y, i, eps, lim)
		pMx = max(pMx, math.Abs(cur))
	}

	fmt.Println("Спектральный радиус матрицы =", pMx)

	// QR разложение. Нахождение собственных значений

	mx = matrix.FromSlice(3, 3, []float64{
		2, -4, 5,
		-5, -2, -3,
		1, -8, -3,
	})

	A := mx
	for range lim {
		Q, R := A.QR()
		A = R.Mul(Q)

		cureps := 0.0
		for i := 1; i < A.M(); i++ {
			cureps += A[i][0] * A[i][0]
		}

		if math.Sqrt(cureps) <= eps {
			break
		}
	}

	fmt.Println()
	fmt.Println("QR алгоритм:")
	fmt.Println(A)

	e := -A[1][1] - A[2][2]
	f := -A[1][2]*A[2][1] + A[1][1]*A[2][2]
	D := e*e - 4*f

	x1 := A[0][0]
	x2 := complex(-e/2, +math.Sqrt(-D)/2)
	x3 := complex(-e/2, -math.Sqrt(-D)/2)

	fmt.Println("Собственные значения")
	fmt.Printf("x1 = %5.3f\n", x1)
	fmt.Printf("x2 = %5.3f\n", x2)
	fmt.Printf("x3 = %5.3f\n", x3)
}
