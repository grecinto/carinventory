package vehicle_data_access

// Result of the CRUD operation on the item.
type Result struct {
	ItemIndex int
	Error string		`json:"Error,omitempty"`
}

// GetResult is a Result with data retrieved from backend DB
type GetResult struct {
	ItemIndex int
	Error string		`json:"Error,omitempty"`
	Vehicle Vehicle		`json:"Vehicle,omitempty"`
}

// IsSuccessful returns true if success, false otherwise
func (r Result) IsSuccessful() bool {
	return r.Error == ""
}

// IsFound returns true if found, false otherwise
func (r GetResult) IsFound() bool {
	return r.Error == "" && !r.Vehicle.IsEmpty()
}

// DeleteAllSucceeded true means all succeeded
func DeleteAllSucceeded(results []Result) bool {
	for _,r := range results{
		if !r.IsSuccessful(){
			return false
		}
	}
	return true
}

// DeleteAllFailed returns true if all items in results failed deletion.
func DeleteAllFailed(results []Result) bool {
	for _,r := range results{
		if r.IsSuccessful(){
			return false
		}
	}
	return true
}

// AllNotFound returns true if results are all not found.
func AllNotFound(results []GetResult) bool {
	for _,r := range results{
		if r.IsFound(){
			return false
		}
	}
	return true
}
