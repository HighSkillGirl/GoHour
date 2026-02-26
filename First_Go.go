package main

import "fmt"

var test_var int  // в рамках пакета можно иметь неиспользуемые переменные
var array [10]int // создание массива. Длина считается частью уникальности - приравнять массивы разной длины не получится
var b = 8

func main() {
	main2("Anastasia")
	name := "Vera" // такой синтаксис только внутри функции!
	main2(name)
	// var test2 int // внутри функции нельзя оставлять неиспользуемые переменные!
	multiplicationTablePrint()
	rangeString("Hello")
}

func main2(name string) {
	a := 10

	if a > b {
		fmt.Println("Hello,", name) // 2 параметра! и разделены пробелом
	} else if (a*2 + 1) > b {
		fmt.Println("Hello, username")
	} else {
		fmt.Println("Goodbye!")
		return
	}

	switch b {
	case 1, 2, 3:
		fmt.Println("little digit", b) // по умолчанию выходит из цикла по совпадению
	case 4, 5, 6:
		fmt.Println("middle digit", b)
	case 7, 8, 9:
		fmt.Println("upper middle digit", b)
		fallthrough // оператор перекидывает на команду для следующего кейса независимо от совпадения
	case 10:
		fmt.Println("the TOP", b)
		b += 1
	}
}
