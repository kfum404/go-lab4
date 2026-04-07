package main

import (
	"fmt"
	"://github.com/kfum404/go-lab4/edit/main/app/main.go"
)

func main() {
	// Создаем объект через твой конструктор из модуля transport
	t := transport.NewSpecial(1, 200, 4, 1.5, "Sportcar")
	
	fmt.Printf("Машина: %s, Скорость: %.2f\n", t.GetKind(), t.GetSpeed())
	t.StartUP()
	t.Gas()
}
