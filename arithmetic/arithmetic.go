package arithmetic

import "fmt"

func ArithmeticOperate(a float64, b string, c float64) string {
	if b == "+" {
		return fmt.Sprint(a + c)
	}
	if b == "-" {
		return fmt.Sprint(a - c)
	}
	if b == "*" {
		return fmt.Sprint(a * c)
	}
	return fmt.Sprint(a / c)
}
