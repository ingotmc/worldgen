package octavenoise

import (
	"github.com/ingotmc/worldgen/noise"
	"math"
	"math/rand"
)

const defaultFreq = 3200.0

type Noise struct {
	octaves   []noise.Noise
	frequency float64
	scalingFactor   float64
	clamp     float64
}

type Option func(noise Noise) Noise

func WithFrequency(freq float64) Option {
	return func(noise Noise) Noise {
		noise.frequency = freq
		return noise
	}
}

func WithScalingFactor(fact float64) Option{
	return func(noise Noise) Noise {
		noise.scalingFactor = fact
		return noise
	}
}

func New(noiseFactory noise.Factory, n int, opts ...Option) noise.Noise {
	octaves := make([]noise.Noise, n)
	for i := range octaves {
		octaves[i] = noiseFactory(rand.Int63())
	}
	res := Noise{
		octaves:   octaves,
		frequency: defaultFreq,
		scalingFactor: 0.5,
	}
	for _, opt := range opts {
		res = opt(res)
	}
	res.clamp = 1.0 / (1.0 - (1.0 / math.Pow(1.0/res.scalingFactor, float64(n))))
	return res
}

func From(ns []noise.Noise, opts ...Option) noise.Noise {
	res := Noise{
		octaves:   ns,
		frequency: defaultFreq,
		clamp:     1.0 / (1.0 - (1.0 / math.Pow(2, float64(len(ns))))),
	}
	for _, opt := range opts {
		res = opt(res)
	}
	return res
}

func (o Noise) Sample(x, z float64) float64 {
	amplFreq := 1.0
	res := 0.0
	for _, noise := range o.octaves {
		inX := x / (amplFreq * o.frequency)
		inZ := z / (amplFreq * o.frequency)
		v := noise.Sample(inX, inZ)
		res += amplFreq * v
		amplFreq *= o.scalingFactor
	}
	res *= o.clamp
	return res
}
