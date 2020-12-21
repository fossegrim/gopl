package weightconv

import "fmt"

type Kilogram float64
type Pound float64

// Approximately a kilogram in pounds.
const KilogramP Pound = 2.2

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }
