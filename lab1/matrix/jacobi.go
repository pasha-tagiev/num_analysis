package matrix

import "math"

func upperMax(mx Matrix) (_, _ int) {
	i := 0
	j := 1
	for m := 0; m < mx.M()-1; m++ {
		for n := m + 1; n < mx.N(); n++ {
			max := math.Abs(mx[i][j])
			cur := math.Abs(mx[m][n])
			if max < cur {
				i = m
				j = n
			}
		}
	}
	return i, j
}

func Jacobi(mx Matrix, eps float64, limit int) ([]float64, Matrix) {
	A := mx.Copy()
	V := E(mx.M())
	for range limit {
		i, j := upperMax(A)

		phi := math.Pi / 4

		if diff := A[i][i] - A[j][j]; math.Abs(diff) > epsilon() {
			phi = math.Atan(2*A[i][j]/diff) / 2
		}

		U := E(mx.M())

		sin := math.Sin(phi)
		cos := math.Cos(phi)

		U[i][i] = +cos
		U[i][j] = -sin
		U[j][i] = +sin
		U[j][j] = +cos

		A = U.Transpose().Mul(A).Mul(U)
		V = V.Mul(U)

		cureps := 0.0
		for m := range A.M() - 1 {
			for n := m + 1; n < A.N(); n++ {
				cureps += A[m][n] * A[m][n]
			}
		}
		cureps = math.Sqrt(cureps)

		if cureps < eps {
			break
		}
	}

	return A.Diagonal(), V
}

func PowerIteration(mx Matrix, y Vector, i int, eps float64, lim int) float64 {
	curY := mx.MulByVec(y)
	lambda := curY[i] / y[i]

	for range lim {
		y = curY
		curY = mx.MulByVec(y)

		prevLambda := lambda
		lambda = curY[i] / y[i]

		if math.Abs(lambda-prevLambda) <= eps {
			break
		}
	}

	return lambda
}
