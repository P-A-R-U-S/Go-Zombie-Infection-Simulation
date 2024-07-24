package main

import (
	data "github.com/P-A-R-U-S/Go-Zombie-Infection-Simulation/Data"
)

func main() {
	world := data.NewWorld()

	// generate humans
	for i := 0; i < 100; i++ {
		entity := data.NewEntity(data.Human)
		world.SetToCell(entity)
	}

	zombie := data.NewEntity(data.Zombie)
	world.SetToCell(zombie)

	world.Draw()
}
