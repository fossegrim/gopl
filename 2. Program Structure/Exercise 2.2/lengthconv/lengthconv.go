// Packcage lengthconv performs feet, inch and meter conversions.
package lengthconv

import "fmt"

type Feet float64
type Inch float64
type Meter float64

const (
	// Approximately one meter in feet.
	MeterF Feet = 3.28
	// Approximately one meter in inches.
	MeterI Inch = 39.37
)

func (f Feet) String() string { return fmt.Sprintf("%gft", f) }
func (i Inch) String() string {
	return fmt.Sprintf("%gin", i)
}
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
