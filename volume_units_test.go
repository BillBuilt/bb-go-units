package units

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVolumeSystems(t *testing.T) {
	SI := "metric"
	assert.Equal(t, SI, Liter.System())
	assert.Equal(t, SI, ExaLiter.System())
	assert.Equal(t, SI, PetaLiter.System())
	assert.Equal(t, SI, TeraLiter.System())
	assert.Equal(t, SI, GigaLiter.System())
	assert.Equal(t, SI, MegaLiter.System())
	assert.Equal(t, SI, KiloLiter.System())
	assert.Equal(t, SI, HectoLiter.System())
	assert.Equal(t, SI, DecaLiter.System())
	assert.Equal(t, SI, DeciLiter.System())
	assert.Equal(t, SI, CentiLiter.System())
	assert.Equal(t, SI, MilliLiter.System())
	assert.Equal(t, SI, MicroLiter.System())
	assert.Equal(t, SI, NanoLiter.System())
	assert.Equal(t, SI, PicoLiter.System())
	assert.Equal(t, SI, FemtoLiter.System())
	assert.Equal(t, SI, AttoLiter.System())

	BI := "imperial"
	assert.Equal(t, BI, ImpQuart.System())
	assert.Equal(t, BI, ImpPint.System())
	assert.Equal(t, BI, ImpGallon.System())
	assert.Equal(t, BI, ImpFluidOunce.System())

	US := "us"
	assert.Equal(t, US, Quart.System())
	assert.Equal(t, US, Pint.System())
	assert.Equal(t, US, Gallon.System())
	assert.Equal(t, US, FluidOunce.System())
}
