package noiseoperator

import (
	"github.com/ingotmc/worldgen/noise"
	"math"
)

var Sigmoid noise.NoiseOperatorFunc = func(n noise.Noise) noise.Noise {
	return noise.NoiseFunc(func(x, y float64) float64 {
		v := n.Sample(x, y)
		return 256.0 / (math.Exp(7.0/3.0-v/64.0) + 1)
	})
}

func Clamp(n noise.Noise) noise.Noise {
	return noise.NoiseFunc(func(x, y float64) float64 {
		v := n.Sample(x, y)
		if v < -1 {
			return -1
		}
		if v > 1 {
			return 1
		}
		return v
	})
}

func ScaleHiLo(hi, lo float64) noise.NoiseOperatorFunc {
	return func(n noise.Noise) noise.Noise {
		return noise.NoiseFunc(func(x, y float64) float64 {
			v := n.Sample(x, y)
			if v > 0 {
				return v * hi
			}
			return v * lo
		})
	}
}

func Offset(off float64) noise.NoiseOperatorFunc {
	return func(n noise.Noise) noise.Noise {
		return noise.NoiseFunc(func(x, y float64) float64 {
			v := n.Sample(x, y)
			return v + off
		})
	}
}

func Abs(n noise.Noise) noise.Noise {
	return noise.NoiseFunc(func(x, y float64) float64 {
		return math.Abs(n.Sample(x, y))
	})
}

func Add(other noise.Noise) noise.NoiseOperatorFunc {
	return func(n noise.Noise) noise.Noise {
		return noise.NoiseFunc(func(x, y float64) float64 {
			return n.Sample(x, y) + other.Sample(x, y)
		})
	}
}