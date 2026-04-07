package main

import (
	"fmt"
	"://github.com"
)

func main() {
	// Создаем объект через твой конструктор из модуля transport
	t := transport.NewSpecial(1, 200, 4, 1.5, "Sportcar")
	
	fmt.Printf("Машина: %s, Скорость: %.2f\n", t.GetKind(), t.GetSpeed())
	t.StartUP()
	t.Gas()
}
