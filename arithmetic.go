package main

var arithmetic map[string]float64

func arithmeticOperate(a float64, b string, c float64) float64 {
	if b == "+" {
		arithmetic[b] = a + c
		return arithmetic[b]
	}
	if b == "-" {
		arithmetic[b] = a - c
		return arithmetic[b]
	}
	if b == "*" {
		arithmetic[b] = a * c
		return arithmetic[b]
	}
	arithmetic[b] = a / c
	return arithmetic[b]
}
