package vehicle_data_access

// NOTE: enumerated values/constant are just for convenience, helper values.

// MakeType enumerated strings
type MakeType string
const(
	Toyota = "Toyota"
	Honda = "Honda"
	MercedesBenc = "Mercedes Benz"
	BMW = "BMW"
	Audi = "Audi"
)

// ModelType enumerated strings
type ModelType string
const(
	CamrySE = "Camry SE"
	CamryLE = "Camry LE"
	CorollaSE = "Corolla SE"
	CorollaLE = "Corolla LE"
)

// ColorType enumerated strings
type ColorType string
const(
	Blue = "Blue"
	AquaMarine = "Aqua Marine"
)

// DoorCountType enum
type DoorCountType int
const(
	FourDoor = iota
	TwoDoor
)

// TrimLevelType
type TrimLevelType string
const(
	Basic = "Basic Trim Level"
	AllLeather = "All Leather"
)

// Vehicle structure
type Vehicle struct {
	VIN string 				`validate:"required" json:"VIN,omitempty"`

	Make string 			`validate:"required" json:"Make,omitempty"`
	Model string 			`validate:"required" json:"Model,omitempty"`
	Color string 			`validate:"required" json:"Color,omitempty"`

	BuyPrice float32 		`validate:"required" json:"BuyPrice,omitempty"`
	MSRP float32 			`validate:"required" json:"MSRP,omitempty"`

	DoorCount int			`json:"DoorCount,omitempty"`
	TrimLevel string		`json:"TrimLevel,omitempty"`
}

// Returns true if Vehicle struct is empty or considered uninitialized
func (v Vehicle)IsEmpty() bool{
	return v.VIN == ""
}
