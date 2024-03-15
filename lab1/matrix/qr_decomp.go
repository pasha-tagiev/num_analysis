package matrix

import "math"

func (mx Matrix) QR() (Matrix, Matrix) {
	Q := E(mx.M())
	R := mx.Copy()
	for i := range R.M() - 1 {
		v := New(R.M(), 1)
		sum := 0.0
		for m := i; m < R.M(); m++ {
			sum += R[m][i] * R[m][i]
		}

		a := R[i][i]
		v[i][0] = a + math.Copysign(1, a)*math.Sqrt(sum)
		for m := i + 1; m < v.M(); m++ {
			v[m][0] = R[m][i]
		}

		num := v.Transpose().Mul(v)
		H := v.Mul(v.Transpose())
		H = H.MulByNum(2 / num[0][0])

		H = E(mx.M()).Sub(H)

		Q = Q.Mul(H)
		R = H.Mul(R)
	}

	return Q, R
}
