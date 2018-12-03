package isa

import (
	"math"
)

type Atmosphere struct {
	Altitude         float64 // [m]
	Temperature      float64 // [deg K]
	Pressure         float64 // [Pa]
	Density          float64 // [kg/m3]
	SpeedOfSound     float64 // [m/s]
	DynamicViscosity float64 // [Pa s]
}

type Calculator struct {
	PressureSeaLevel    float64 // Pressure at sea level [Pa]
	TemperatureSeaLevel float64 // Temperature at sea level [deg K]
}

const (
	airMolWeight  = 28.9644 // Molecular weight of air [g/Mol]
	densitySL     = 1.225   // Density at sea level [kg/m3]
	pressureSL    = 101325  // Pressure at sea level [Pa]
	temperatureSL = 288.15  // Temperature at sea level [deg K]
	gamma         = 1.4
	gravity       = 9.80665 // Acceleration of gravity [m/s2]
	RGas          = 8.31432 // Gas constant [kg/Mol/K]
	R             = 287.053 //
	gMR           = gravity * airMolWeight / RGas
)

func NewCalculator() Calculator {
	return Calculator{
		PressureSeaLevel:    pressureSL,
		TemperatureSeaLevel: temperatureSL,
	}
}

var (
	altitudes    = []float64{0, 11000, 20000, 32000, 47000, 51000, 71000, 84852}
	pressureRels = []float64{1, 2.23361105092158e-1, 5.403295010784876e-2, 8.566678359291667e-3, 1.0945601337771144e-3, 6.606353132858367e-4, 3.904683373343926e-5, 3.6850095235747942e-6}
	temperatures = []float64{288.15, 216.65, 216.65, 228.65, 270.65, 270.65, 214.65, 186.946}
	tempGrads    = []float64{-6.5, 0, 1, 2.8, 0, -2.8, -2, 0}
)

func (c Calculator) Calculate(altitude float64) Atmosphere {
	if altitude < -2000 || altitude > 86000 {
		panic("altitude must be between -2000 and 86000 meters")
	}

	i := 0
	if altitude > 0 {
		for ; altitude > altitudes[i+1]; i++ {
		}
	}

	lBaseTemp := temperatures[i]
	tempGrad := tempGrads[i] / 1000
	lPressureRelBase := pressureRels[i]
	lDeltaAltitude := altitude - altitudes[i]
	temperature := lBaseTemp + tempGrad*lDeltaAltitude

	var lPressureRelative float64
	if tempGrad == 0 {
		lPressureRelative = lPressureRelBase * math.Exp(-gMR*lDeltaAltitude/1000/lBaseTemp)
	} else {
		lPressureRelative = lPressureRelBase * math.Pow(lBaseTemp/temperature, gMR/tempGrad/1000)
	}

	temperature += c.TemperatureSeaLevel - temperatureSL

	return Atmosphere{
		Altitude:         altitude,
		Temperature:      temperature,
		Pressure:         lPressureRelative * c.PressureSeaLevel,
		Density:          densitySL * lPressureRelative * temperatureSL / temperature,
		SpeedOfSound:     math.Sqrt(gamma * R * temperature),
		DynamicViscosity: 1.512041288 * math.Pow(temperature, 1.5) / (temperature + 120) / 1000000.0,
	}
}
