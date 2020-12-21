package lengthconv

// FToI converts a feet length to inches.
func FToI(f Feet) Inch {
	return MToI(FToM(f))
}

// FToM converts a feet length to meters.
func FToM(f Feet) Meter {
	return Meter(f / MeterF)
}

// IToF converts a inch length to feet.
func IToF(i Inch) Feet {
	return MToF(IToM(i))
}

// IToM converts a inch length to feet.
func IToM(i Inch) Meter {
	return Meter(i / MeterI)
}

// MToF converts a meter length to feet.
func MToF(m Meter) Feet {
	return Feet(m) * MeterF
}

// MToI converts a meter length to inches.
func MToI(m Meter) Inch {
	return Inch(m) * MeterI
}
