package main

import "archh/Vanish/service"

func main() {
	pist := &service.Pistol{Ammo: 3}
	service.Fire(pist)
}
