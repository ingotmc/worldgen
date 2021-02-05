package ridge

import (
	"github.com/ingotmc/worldgen/noise"
	"math"
)

type Noise struct {
	source noise.Noise
}

func (n Noise) Apply(other noise.Noise) noise.Noise {
	return noise.NoiseFunc(func(x, y float64) float64 {
		return other.Sample(x, y) + (1 - math.Abs(n.source.Sample(x,y))) * 80.0
	})
}

func New(ns noise.Noise) noise.NoiseOperator {
	return Noise{source: ns}
}

