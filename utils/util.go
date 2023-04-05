package utils

func ConcatenateRecommendedRestaurants(highPriorityRestaurants []string, lowPriorityRestaurants []string) []string {
	var concatenatedResult []string
	restaurants := make(map[string]bool)
	for _, restaurantId := range highPriorityRestaurants {
		_, found := restaurants[restaurantId]
		if !found {
			concatenatedResult = append(concatenatedResult, restaurantId)
			restaurants[restaurantId] = true
		}
	}
	for _, restaurantId := range lowPriorityRestaurants {
		_, found := restaurants[restaurantId]
		if !found {
			concatenatedResult = append(concatenatedResult, restaurantId)
			restaurants[restaurantId] = true
		}
	}
	return concatenatedResult
}
