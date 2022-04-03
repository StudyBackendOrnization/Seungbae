package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	answer += compute(data)
	println("결과 값 : ")
	calculator()
}

func compute(data string) float64 {
	return 1.0
}

//func operate(x float64, sign string, y float64) float64 {
//	return arithmeticOperate(x, sign, y)
//}

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
