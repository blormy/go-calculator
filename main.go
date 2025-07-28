// Простой консольный калькулятор
package main

import (
	"fmt"
	"os"
	"strings"
)

// Запускает цикл калькулятора. Запрашивает у пользователя числа и операцию.
// Выводит результат. Спрашивает о продолжении работы
func main() {
	var again string
	for {
		a, b, op := getData()
		res := calculator(a, b, op)
		fmt.Println("Результат: ", res)

		fmt.Println("Продолжить работу с калькулятором? (y/n): ")
		fmt.Scanln(&again)
		if strings.ToLower(again) != "y" {
			fmt.Println("Работа завершена")
			break
		}
	}
}

// getData запрашивает у пользователя два числа и операцию.
// Возвращает введенные значения: первое число, второе число, операцию над числами.
// В случае неправильного ввода завершает программу
func getData() (float64, float64, string) {
	var a, b float64
	var op string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Неккоректный ввод")
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Неккоректный ввод")
		os.Exit(1)
	}

	fmt.Print("Введите нужную операцию (+, -, /, *): ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println("Неккоректный ввод")
		os.Exit(1)
	}

	return a, b, op
}

// calculator выполняет операцию над двумя числами.
// В случае деления на ноль или неизвестной операции завершает программу с ошибкой.
func calculator(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: на ноль делить нельзя")
			os.Exit(1)
		}
		return a / b
	default:
		fmt.Println("Ошибка: неизвестная операция %s", op)
		os.Exit(1)
		return 0
	}
}
