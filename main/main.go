package main

import (
	"bufio"
	"fmt"
	"main/arithmetic"
	"math"
	"os"
	"strconv"
	"strings"
)

var answer float64

func main() {
	answer = 0
	fmt.Println("---------------------")
	fmt.Println("        계산기       ")
	fmt.Println("---------------------")

	calculator()
}

func calculator() {
	data, err := DataScan()

	if err != nil {
		println(err.Error())
		calculator()
		return
	}

	data = fmt.Sprint(answer) + data
	answer = compute(data)
	fmt.Printf("결과 값 : %f\n", tofixed(answer, 5))
	calculator()
}

func compute(data string) float64 {
	if strings.ContainsAny(data, "*") || strings.ContainsAny(data, "/") {
		x := strings.IndexAny(data, "*")
		y := strings.IndexAny(data, "/")

		if x > y {
			if y != -1 {
				back := dataFindPlus(data, y+1)
				front := dataFindMinus(data, y-1)
				data = data[0:front] + operate(data[front:y], string(data[y]), data[y+1:back]) + data[back:]
				return compute(data)
			}
			back := dataFindPlus(data, x+1)
			front := dataFindMinus(data, x-1)
			data = data[0:front] + operate(data[front:x], string(data[x]), data[x+1:back]) + data[back:]
			return compute(data)
		}

		if x < y {
			if x != -1 {
				back := dataFindPlus(data, x+1)
				front := dataFindMinus(data, x-1)
				data = data[0:front] + operate(data[front:x], string(data[x]), data[x+1:back]) + data[back:]
				return compute(data)
			}
			back := dataFindPlus(data, y+1)
			front := dataFindMinus(data, y-1)
			data = data[0:front] + operate(data[front:y], string(data[y]), data[y+1:back]) + data[back:]
			return compute(data)
		}
	}

	if strings.ContainsAny(data, "+") || strings.ContainsAny(data, "-") {
		x := strings.IndexAny(data, "+")
		y := strings.IndexAny(data, "-")

		if x > y {
			if y != -1 {
				back := dataFindPlus(data, y+1)
				front := dataFindMinus(data, y-1)
				data = data[0:front] + operate(data[front:y], string(data[y]), data[y+1:back]) + data[back:]
				return compute(data)
			}
			back := dataFindPlus(data, x+1)
			front := dataFindMinus(data, x-1)
			data = data[0:front] + operate(data[front:x], string(data[x]), data[x+1:back]) + data[back:]
			return compute(data)
		}

		if x < y {
			if x != -1 {
				back := dataFindPlus(data, x+1)
				front := dataFindMinus(data, x-1)
				data = data[0:front] + operate(data[front:x], string(data[x]), data[x+1:back]) + data[back:]
				return compute(data)
			}
			back := dataFindPlus(data, y+1)
			front := dataFindMinus(data, y-1)
			data = data[0:front] + operate(data[front:y], string(data[y]), data[y+1:back]) + data[back:]
			return compute(data)
		}
	}
	result, _ := strconv.ParseFloat(data, 64)
	return result
}

func dataFindPlus(data string, index int) int {
	if index == len(data) {
		return index
	}
	if '0' <= data[index] && data[index] <= '9' {
		return dataFindPlus(data, index+1)
	}
	return index
}

func dataFindMinus(data string, index int) int {
	if index == -1 {
		return index + 1
	}
	if '0' <= data[index] && data[index] <= '9' {
		return dataFindMinus(data, index-1)
	}
	return index + 1
}

func operate(x string, sign string, y string) string {
	xResult, _ := strconv.ParseFloat(x, 64)
	yResult, _ := strconv.ParseFloat(y, 64)
	return arithmetic.ArithmeticOperate(xResult, sign, yResult)
}

func DataScan() (string, error) {
	s := bufio.NewScanner(os.Stdin)

	fmt.Printf("데이터를 입력해주세요 :%f", tofixed(answer, 5))
	s.Scan()

	input := strings.ToLower(s.Text())

	if strings.EqualFold(input, "exit") {
		os.Exit(3)
	}

	if strings.EqualFold(input, "c") {
		answer = 0
		return "", fmt.Errorf("결과값 : 0")
	}

	if err := checkDataRight(input, 0); err != nil {
		return "", err
	}

	return input, nil
}

func tofixed(answer float64, count int) float64 {
	fix := math.Pow(10, float64(count))
	return float64(round(answer*fix)) / fix
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func checkDataRight(data string, count int) error {
	chars := []rune(data)
	if count == len(chars) {
		return nil
	}
	if '0' <= chars[count] && chars[count] <= '9' || chars[count] == '*' || chars[count] == '-' || chars[count] == '+' || chars[count] == '/' {
		return checkDataRight(data, count+1)
	}
	return fmt.Errorf("다시 입력해주세요.\n")
}
