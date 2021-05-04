package vehicle_data_access

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
	VIN string 				`validate:"required"`

	Make MakeType 			`validate:"required"`
	Model ModelType 		`validate:"required"`
	Color ColorType 		`validate:"required"`

	BuyPrice float32 		`validate:"required"`
	MSRP float32 			`validate:"required"`

	DoorCount DoorCountType
	TrimLevel TrimLevelType
}

// Returns true if Vehicle struct is empty or considered uninitialized
func (v Vehicle)IsEmpty() bool{
	return v.VIN == ""
}
