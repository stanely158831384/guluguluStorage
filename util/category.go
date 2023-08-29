package util

const (
	produces = "produces"
	fruits = "fruits"
	meat = "meat"
	condiments = "condiments"
	beverages = "beverages"
	snacks = "snacks"
	seasonings = "seasonings"
	breadgrains = "breadgrains"
	dairy = "dairy"
	frozenfoods = "frozenfoods"
	cannedfoods = "cannedfoods"
)
func CategoryDetector(input string)bool{
	switch input {
	case produces, fruits, meat, condiments, beverages, snacks, seasonings, breadgrains, dairy, frozenfoods, cannedfoods:
		return true
	}
	return false
}