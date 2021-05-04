package vehicle_data_access

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var cars []Vehicle = []Vehicle{
	Vehicle{
		VIN: "1234",
		Make: Toyota,
		Model: CamryLE,
		Color: AquaMarine,
		BuyPrice: 20000,
		MSRP: 25000,
		DoorCount: FourDoor,
		TrimLevel: Basic,
	},
	Vehicle{
		VIN: "1235",
		Make: Toyota,
		Model: CamrySE,
		Color: Blue,
		BuyPrice: 20001,
		MSRP: 25001,
		DoorCount: FourDoor,
		TrimLevel: Basic,
	},
}

// hacky unit tests, hehe, not much time. :)

func TestVehicleDataAccessSet(t *testing.T){
	var da VehicleDataAccess = NewMockAccess()
	if len(Validate(cars...)) > 0{
		t.Error("cars validation failed.")
		return
	}
	results := da.Set(cars...)
	assert.Equal(t, 0, len(results), "Set should return 0 results")
}

func TestVehicleDataAccessDelete(t *testing.T){
	var da VehicleDataAccess = NewMockAccess()
	if len(Validate(cars...)) > 0{
		t.Error("cars validation failed.")
		return
	}
	results := da.Set(cars...)
	assert.Equal(t, 0, len(results), "Set should return 0 results")
	results = da.Delete([]string{"1234"}...)
	assert.True(t, results[0].IsSuccessful(), "Failed to delete car with VIN 1234")
}

func TestVehicleDataAccessGet(t *testing.T){
	var da VehicleDataAccess = NewMockAccess()
	if len(Validate(cars...)) > 0{
		t.Error("cars validation failed.")
		return
	}
	results := da.Set(cars...)
	assert.Equal(t, 0, len(results), "Set should return 0 results")

	getResults := da.Get([]string{"1234"}...)
	assert.True(t, getResults[0].IsFound(), "Failed to get car with VIN 1234")
}

func TestVehicleDataAccessSetDeleteGet(t *testing.T){
	var da VehicleDataAccess = NewMockAccess()
	if len(Validate(cars...)) > 0{
		t.Error("cars validation failed.")
		return
	}
	results := da.Set(cars...)
	assert.Equal(t, 0, len(results), "Set should return 0 results")

	da.Delete([]string{"1234"}...)

	getResults := da.Get([]string{"1234"}...)
	assert.False(t, getResults[0].IsFound(), "Able to get deleted car with VIN 1234")
}
