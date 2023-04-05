package domain

type CuisineTracking struct {
	CuisineType Cuisine
	NoOfOrders  int
}

type CostTracking struct {
	CostTrackingType int
	NoOfOrders       int
}

type User struct {
	TrackedCuisines []CuisineTracking
	TrackedCosts    []CostTracking
}

type UserCuisinePreference struct {
	PrimaryCuisine   Cuisine
	SecondaryCuisine []Cuisine
}

type UserCostPreference struct {
	PrimaryCostbracket   int
	SecondaryCostbracket []int
}
type UserPreference struct {
	UserCuisinePreference UserCuisinePreference
	UserCostPreference    UserCostPreference
}
