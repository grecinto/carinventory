package vehicle_data_access

import (
	"sync"
	"fmt"
)

type mockAccess struct{
	cache map[string]Vehicle
	mutex *sync.Mutex
}

// NewMockDataAccess returns a Vehicle Data Access with caching in-memory.
// Designed for doing mockup and unit testing.
func NewMockAccess() VehicleDataAccess{
	return &mockAccess{
		cache: map[string]Vehicle{},
		mutex: &sync.Mutex{},
	}
}

func (ma *mockAccess)Set(vehicles ...Vehicle) []Result {
	ma.mutex.Lock()
	defer ma.mutex.Unlock()
	for _,v := range vehicles {
		ma.cache[v.VIN] = v
	}
	return nil
}

func (ma mockAccess)Get(vins ...string) []GetResult{
	ma.mutex.Lock()
	defer ma.mutex.Unlock()
	var results []GetResult
	for i,vin := range vins {
		v := ma.cache[vin]
		r := GetResult{}
		if v.IsEmpty() {
			r.Error = fmt.Sprintf("Didn't find item with VIN '%s'", vin)
		} else {
			r.Vehicle = v
		}
		r.ItemIndex = i
		results = append(results, r)
	}
	return results
}

func (ma *mockAccess)Delete(vins ...string) []Result{
	ma.mutex.Lock()
	defer ma.mutex.Unlock()
	var results []Result
	for i,vin := range vins {
		v := ma.cache[vin]
		r := Result{
			ItemIndex: i,
		}
		if v.IsEmpty() {
			r.Error = fmt.Sprintf("Can't delete item with VIN '%s', 'not found.", vin)
		}
		results = append(results, r)
	}

	// delete the items now that we've determined their existence.
	for _,vin := range vins {
		delete(ma.cache,vin)
	}
	return results
}
