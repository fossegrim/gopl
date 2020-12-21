package weightconv

// KToP converts a kilogram weight to pounds.
func KToP(kg Kilogram) Pound {
	return Pound(kg) * KilogramP
}

// PToK converts a pound weight to kilograms.
func PToK(lb Pound) Kilogram {
	return Kilogram(lb / KilogramP)
}
