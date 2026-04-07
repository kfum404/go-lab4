package main

import (
	"fmt"
	"github.com/kfum404/go-lab4/transport"
)

func main() {
	// Создаем несколько машин, как в лабе 0.5
	trans1 := transport.NewSpecial(67, 160, 6, 1.25, "Medicine")
	trans2 := transport.NewSpecial(52, 220, 4, 1, "Police")
	trans3 := transport.NewSpecial(23, 420, 4, 0.8, "Supercar")

	// Выводим инфу по всем
	fmt.Printf("ID: %d, Скорость: %.2f км/ч, Тип: %s\n", trans1.Id, trans1.Speed, trans1.Kind)
	fmt.Printf("ID: %d, Скорость: %.2f км/ч, Тип: %s\n", trans2.Id, trans2.Speed, trans2.Kind)
	fmt.Printf("ID: %d, Скорость: %.2f км/ч, Тип: %s\n", trans3.Id, trans3.Speed, trans3.Kind)

	// Вызываем методы
	trans1.Voice()
	trans3.Voice()
	trans1.Gas()
	trans2.Autopilot()
	trans3.StartUP()
	trans3.TurnOFF()

	fmt.Printf("%s Speed: %.2f км/ч\n", trans1.GetKind(), trans1.GetSpeed())
}

