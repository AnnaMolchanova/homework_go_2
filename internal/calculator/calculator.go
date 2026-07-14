package calculator

import "errors"

// Блок calculator — объединяющая задача.
// Здесь нужно закрепить:
// int, string operation, switch, if, error,
// целочисленное деление и остаток от деления.

// Calculate выполняет арифметическую операцию над двумя int-числами.
// Деление и остаток от деления должны безопасно обрабатывать нулевой делитель.
// Неизвестная операция должна возвращать ошибку.
func Calculate(a, b int, operation string) (int, error) {
	if operation == "+" {
		return a + b, nil
	} else if operation == "-" {
		return a - b, nil
	} else if operation == "*" {
		return a * b, nil
	} else if operation == "/" {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil

	} else if operation == "%" {
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a % b, nil

	} else {
		return 0, errors.New("invalid operation")
	}

}
