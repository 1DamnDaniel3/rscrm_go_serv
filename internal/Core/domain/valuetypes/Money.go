package valuetypes

import "math"

type Money int64

func (m Money) Add(o Money) Money { return m + o }
func (m Money) Sub(o Money) Money { return m - o }
func (m Money) MulPercent(p float64) Money {
	return Money(math.Round(float64(m) * p / 100))
}
