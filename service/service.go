package service

import (
	"fmt"
	"math/rand"
	"time"
)

type Pistol struct {
	Ammo int
}

type Release interface {
	Shoot()
	Reload()
}

func (p *Pistol) Reload() {
	fmt.Println("Reload...")
	p.Ammo = 6
	fmt.Println("Reload done✅")
}

func (p *Pistol) Shoot() {
	rand.Seed(time.Now().UnixNano())
	var wearBroke bool = true
	for wearBroke {
		if p.Ammo == 0 {
			p.Reload()
			wearBroke = false
		}
		for p.Ammo != 0 {
			randomNum := rand.Intn(101)
			if randomNum%2 == 0 {
				fmt.Println("Bang✨")
				p.Ammo--
			} else {
				fmt.Println("No Bang😁")
				p.Ammo--
			}
		}
	}
}

func Fire(r Release) {
	r.Shoot()
}
