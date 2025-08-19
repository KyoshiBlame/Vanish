package main

import "Vanish/service"

func main() {
	pist := &service.Pistol{Ammo: 3}
	service.Fire(pist)
	//too
}
