package units

var (
	Volume = UnitOptionQuantity("volume")

	// metric
	Liter      = NewUnit("liter", "l", Volume, SI, UnitOptionAliases("litre"))
	ExaLiter   = Exa(Liter)
	PetaLiter  = Peta(Liter)
	TeraLiter  = Tera(Liter)
	GigaLiter  = Giga(Liter)
	MegaLiter  = Mega(Liter)
	KiloLiter  = Kilo(Liter)
	HectoLiter = Hecto(Liter)
	DecaLiter  = Deca(Liter)
	DeciLiter  = Deci(Liter)
	CentiLiter = Centi(Liter)
	MilliLiter = Milli(Liter)
	MicroLiter = Micro(Liter)
	NanoLiter  = Nano(Liter)
	PicoLiter  = Pico(Liter)
	FemtoLiter = Femto(Liter)
	AttoLiter  = Atto(Liter)

	// imperial
	ImpQuart      = NewUnit("imperial quart", "imp qt", Volume, BI)
	ImpPint       = NewUnit("imperial pint", "imp pt", Volume, BI)
	ImpGallon     = NewUnit("imperial gallon", "imp gal", Volume, BI)
	ImpFluidOunce = NewUnit("imperial fluid ounce", "imp fl oz", Volume, BI)
	ImpTeaspoon   = NewUnit("imperial teaspoon", "imp tsp", Volume, BI, UnitOptionAliases("t"))
	ImpTablespoon = NewUnit("imperial tablespoon", "imp tbsp", Volume, BI, UnitOptionAliases("T"))
	ImpCup        = NewUnit("imperial cup", "imp c", Volume, BI)

	// US
	Quart      = NewUnit("quart", "qt", Volume, US)
	Pint       = NewUnit("pint", "pt", Volume, US)
	Gallon     = NewUnit("gallon", "gal", Volume, US)
	FluidOunce = NewUnit("fluid ounce", "fl oz", Volume, US, UnitOptionAliases("floz"))
	Teaspoon   = NewUnit("teaspoon", "tsp", Volume, US, UnitOptionAliases("t"))
	Tablespoon = NewUnit("tablespoon", "tbsp", Volume, US, UnitOptionAliases("T"))
	Cup        = NewUnit("cup", "c", Volume, US, UnitOptionAliases("C"))
)

func init() {
	NewRatioConversion(ImpTeaspoon, MilliLiter, 5.919388)
	NewRatioConversion(ImpTablespoon, MilliLiter, 17.75816)
	NewRatioConversion(ImpFluidOunce, MilliLiter, 28.4130625)
	NewRatioConversion(ImpCup, MilliLiter, 284.1306)
	NewRatioConversion(ImpPint, MilliLiter, 568.26125)
	NewRatioConversion(ImpQuart, Liter, 1.1365225)
	NewRatioConversion(ImpGallon, Liter, 4.54609)

	NewRatioConversion(Teaspoon, MilliLiter, 4.928922)
	NewRatioConversion(Tablespoon, MilliLiter, 14.78677)
	NewRatioConversion(FluidOunce, MilliLiter, 29.5735295625)
	NewRatioConversion(Cup, MilliLiter, 236.5882365)
	NewRatioConversion(Pint, MilliLiter, 473.176473)
	NewRatioConversion(Quart, Liter, 0.946352946)
	NewRatioConversion(Gallon, Liter, 3.785411784)
}
