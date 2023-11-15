package rpg

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	Name       string
	BaseAttack int
	Defence    int
	CritDamage int
	HP         int
	Weapon     Weapon
}

type Weapon struct {
	Name   string
	Attack int
}

// counts attack damage
func (h Hero) CountDamage() int {
	damage := h.BaseAttack + h.Weapon.Attack
	if rand.Intn(100)-1 > 50 {
		damage += h.CritDamage
	}
	return damage
}

// returns remaining HP
func (h1 Hero) isAttackedBy(h2 Hero) int {
	totalDmg := h2.CountDamage() - h1.Defence
	if totalDmg <= 0 {
		return h1.HP
	} else {
		return h1.HP - totalDmg
	}
}

func Battle(hAtk Hero, hDef Hero) {
	remainingHP := hDef.isAttackedBy(hAtk)

	fmt.Printf("%s attacks %s...\n", hAtk.Name, hDef.Name)
	time.Sleep(2 * time.Second)

	if remainingHP <= 0 {
		fmt.Printf("%s defeats %s!", hAtk.Name, hDef.Name)
	} else if remainingHP == hDef.HP {
		fmt.Printf("%s fully defends %s's attack!", hDef.Name, hAtk.Name)
	} else {
		fmt.Printf("%s defends %s's attack with %d HP left!", hDef.Name, hAtk.Name, remainingHP)
	}
}
