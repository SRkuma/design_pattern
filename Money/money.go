package Money

type Money int64

func Dollars(value int64) Money {
	return Money(value)
}