package main

import "fmt"

type fuel struct {
	name string
	literPerMill int
}

type Car struct {
	Type fuel
	Fuel int
}

func (c *Car) CalcMaxDistance() int {
	maxDistance := c.Fuel/c.Type.literPerMill
	return maxDistance
}


func main() {
	fuelType1 := fuel{"fuel1", 15}
	car1 := Car{Type: fuelType1, Fuel: 15}	
	fmt.Println(car1.CalcMaxDistance())
}