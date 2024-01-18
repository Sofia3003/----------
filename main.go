package main

import (
	structure "app/pack"
	"fmt"
)

func main() {
	cars := make([]structure.Automobile, 0, 4) // Создание среза структур
	car1 := structure.Create_automobile("bmw", 2021)
	car2 := structure.Create_automobile("ford", 2019)
	car3 := structure.Create_automobile("volkswagen", 2020)
	fmt.Println(cars)
	fmt.Println(structure.TryAdd(&cars, *car1))
	fmt.Println(structure.TryAdd(&cars, *car2))
	fmt.Println(structure.TryAdd(&cars, *car3))
	fmt.Println(cars)
	fmt.Println(structure.Average_summ(cars))
}
