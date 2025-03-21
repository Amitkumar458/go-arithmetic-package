package arithmetic

import "fmt"

func Add(numbers ...float64) float64 {
	var sum float64 = 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func Substract(a float64, b float64) float64 {
	return a - b
}

func Multiply(numbers ...float64) (float64, error) {
	var mul float64 = 1.0
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}
	for _, value := range numbers {
		mul = mul * value
	}
	return mul, nil
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}
	return a / b, nil
}
