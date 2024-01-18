package structure

import (
	"errors"
	"fmt"
)

type Automobile struct {
	brand               string
	year_of_manufacture int
	speed               int
}

func Create_automobile(brand_name string, year_of_manufacture int) *Automobile {
	car := new(Automobile)
	car.brand = brand_name
	car.set_year_of_manufacture(year_of_manufacture)
	car.set_speed(brand_name, year_of_manufacture)
	return car
}

func (a *Automobile) set_year_of_manufacture(year_of_manufacture int) error {
	if year_of_manufacture < 1900 {
		err := errors.New("data exceeds the limits")
		fmt.Println("", err)
		return err
	}
	a.year_of_manufacture = year_of_manufacture
	return nil
}

func (a *Automobile) set_speed(brand_name string, year_of_manufacture int) error {
	if brand_name == "bmw" && year_of_manufacture == 2021 {
		a.speed = 140
		return nil
	}
	if brand_name == "ford" && year_of_manufacture == 2019 {
		a.speed = 200
		return nil
	}
	if brand_name == "volkswagen" && year_of_manufacture == 2020 {
		a.speed = 198
		return nil
	}
	err := errors.New("this auto is missing from the catalog")
	a.speed = 0
	fmt.Println("", err)
	return err
}

func TryAdd(slices *[]Automobile, car Automobile) bool {
	slices_value := *slices
	for i := 0; i != len(slices_value); i++ {
		if slices_value[i] == car {
			return false
		}
	}
	*slices = append(*slices, car)
	return true
}

func Average_summ(cars []Automobile) float64 {
	k := len(cars)
	sum := 0.0
	for i := 0; i != k; i++ {
		sum = sum + float64(cars[i].speed)
	}
	return sum / float64(k)
}
