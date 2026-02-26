package main

import "fmt"

func multiplicationTablePrint() {
	fmt.Println("таблица умножения:")
OutherLoop: // метка для выхода из цикла, к которому приклеена метка
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
			if false {
				break OutherLoop
			}

		}
		fmt.Print("\n")
	}
}

func rangeString(stringValue string) {
	for index, value := range stringValue { // итерирование по коллекции - оператор range. Value - копия! элемента, если нужно изменить - обращение по индексу
		fmt.Printf("Index: %d, Value: %d, Letter: %c\n", index, value, value)
	}
}
