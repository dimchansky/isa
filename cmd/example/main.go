package main

import (
	"fmt"
	"math"

	"github.com/dimchansky/isa"
)

func main() {
	c := isa.NewCalculator()
	ro0 := c.Calculate(0).Density

	fmt.Printf("%5s\t%7s\t%13s\t%11s\t%17s\t%6s\t%12s\n", "H", "T", "Pₕ", "ρₕ", "√(ρ₀/ρₕ)", "a", "η")
	for alt := 0.0; alt <= 80000; alt += 500 {
		a := c.Calculate(alt)
		fmt.Printf("%5.0f m\t%7.2f °C\t%9.2f Pa\t%7.5f kg/m3\t%9.5f\t%6.1f m/s\t %.6g Pa.s\n",
			a.Altitude, a.Temperature-273.15, a.Pressure, a.Density, math.Sqrt(ro0/a.Density), a.SpeedOfSound, a.DynamicViscosity)
	}
}
