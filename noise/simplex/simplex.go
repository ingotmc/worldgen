package simplex

import (
	"github.com/ingotmc/worldgen/noise"
	"github.com/ojrac/opensimplex-go"
)

type Noise struct {
	opensimplex.Noise
}

func New(seed int64) noise.Noise {
	return Noise{opensimplex.New(seed)}
}

func (s Noise) Sample(x, y float64) float64 {
	return s.Eval2(x, y)
}
