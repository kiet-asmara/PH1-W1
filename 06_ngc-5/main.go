package main

import (
	"W1-D1/06_ngc-5/rpg"
	"fmt"
)

type Person struct {
	name string
	age  int
	job  string
}

func main() {
	// soal 1
	// john := Person{
	// 	name: "john",
	// 	age:  45,
	// 	job:  "Gambler",
	// }
	// john.AddYear()
	// john.AddYear()
	// john.AddYear()
	// john.AddYear()
	// john.AddYear()
	// john.GetInfo()

	// soal 2
	// var sliceOfPerson = []Person{
	// 	john,
	// 	{name: "james", age: 34, job: "farmer"},
	// 	{name: "joni", age: 30, job: "doctor"},
	// 	{name: "jane", age: 20, job: "nurse"},
	// }

	// for _, p := range sliceOfPerson {
	// 	p.GetInfo()
	// }

	// soal 3

	firstHero := rpg.Hero{
		Name:       "jojo",
		BaseAttack: 30,
		Defence:    30,
		CritDamage: 30,
		HP:         100,
		Weapon:     rpg.Weapon{Name: "golok", Attack: 10},
	}
	// fmt.Printf("%s attacks with %d damage!\n", firstHero.Name, firstHero.CountDamage())

	// soal 4
	secondHero := rpg.Hero{
		Name:       "zaza",
		BaseAttack: 30,
		Defence:    100,
		CritDamage: 30,
		HP:         100,
		Weapon:     rpg.Weapon{Name: "celurit", Attack: 50},
	}
	rpg.Battle(firstHero, secondHero)
}

func (p Person) GetInfo() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Job:", p.job)
	fmt.Println("")
}

func (p *Person) AddYear() {
	p.age += 1
	if p.age >= 50 {
		p.job = "Retired"
	}
}
