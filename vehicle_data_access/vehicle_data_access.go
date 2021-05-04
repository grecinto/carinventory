package vehicle_data_access

import (
	"github.com/go-playground/validator"
)

// NOTE: the VehicleDataAccess interface (optionally) supports array of items CRUD operations.
// Bulk I/O is typically supported in backend DB, thus, array of items can be an easy optimization, utilizing
// Bulk I/O for example, without depending on load balancers for sprinkling requests across
// App instances. But can be supported too.


// VehicleDataAccess CRUD interface
type VehicleDataAccess interface {
	// Set upserts vehicle(s) to the Backend DB
	Set(vehicles ...Vehicle) []Result
	// Get retrieves vehicle(s) with given VINs from Backend DB
	Get(vins ...string) []GetResult
	// Delete removes vehicle(s) with given VINs from Backend DB
	Delete(vins ...string) []Result
}

// Validate the Vehicles structs per annotations
func Validate(vehicles ...Vehicle) []Result{
	var results []Result
	validate := validator.New()
	for i,v := range vehicles {
		err := validate.Struct(v)
		if err != nil {
			results = append(results, Result{
				ItemIndex: i,
				Error: err.Error(),
			})
		}
	}
	return results
}
