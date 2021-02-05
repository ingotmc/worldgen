package noise

type Factory func (seed int64) Noise

type Noise interface {
	Sample(x, y float64) float64
}

type NoiseFunc func(x, y float64) float64

func (nf NoiseFunc) Sample(x, y float64) float64 {
	return nf(x, y)
}

type NoiseOperator interface {
	Apply(noise Noise) Noise
}

type NoiseOperatorFunc func(Noise) Noise

func (nof NoiseOperatorFunc) Apply(noise Noise) Noise {
	return nof(noise)
}

func Apply(noise Noise, op NoiseOperator, others... NoiseOperator) Noise {
	noise = op.Apply(noise)
	for _, operator := range others {
		noise = operator.Apply(noise)
	}
	return noise
}
