package main

import (
	"fmt"
)

func main() {
	var numInput string
	fmt.Print("Введите число в cc с основанием -10: ")
	fmt.Scan(&numInput)

	result, err := Convert(numInput)
	if err != nil {
		fmt.Println("ОШИБКА")
		return
	}
	fmt.Printf("%s(-10) = %d(10)\n", numInput, result)
}

func Convert(num string) (int, error) {
	if len(num) == 0 {
		return 0, fmt.Errorf("пустая строка")
	}

	result := 0
	lenNum := len(num)
	for i, j := range num {
		if j < '0' || j > '9' {
			return 0, fmt.Errorf("ERROR Неверный формат")
		}
		value := int(j - '0')
		power := lenNum - i - 1
		if power%2 == 0 {
			result += value * pow(10, power)
		} else {
			result -= value * pow(10, power)
		}
	}

	return result, nil
}

func pow(num, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= num
	}
	return result
}
