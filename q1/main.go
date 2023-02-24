package main

import "fmt"

type Vechile interface {
	getName() string
	countWheels() int
}

type car struct {
	name   string
	wheels int
}

type truck struct {
	name   string
	wheels int
	cargo  string
}

type bike struct {
	name         string
	wheels       int
	handle_color string
}

// car methods
func (c car) getName() string {
	return c.name
}

func (c car) countWheels() int {
	return c.wheels
}

// truck methods
func (t truck) getName() string {
	return t.name
}
func (t truck) countWheels() int {
	return t.wheels
}

func (b bike) getName() string {
	return b.name
}
func (b bike) countWheels() int {
	return b.wheels
}
func compareWheels(v1 Vechile, v2 Vechile) {
	if v1.countWheels() > v2.countWheels() {
		fmt.Printf("%s has more wheels than %s", v1.getName(), v2.getName())
	} else {
		fmt.Printf("%s has more wheels than %s", v2.getName(), v1.getName())
	}
	// \n
}

func main() {

	car1 := car{
		name:   "honda",
		wheels: 4,
	}

	truck1 := truck{
		name:   "cargo",
		wheels: 6,
	}

	bike1 := bike{
		name:   "bmx",
		wheels: 2,
	}

	compareWheels(car1, truck1)
	compareWheels(car1, bike1)

}
