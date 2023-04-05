package domain

import "time"

type Restaurant struct {
	RestaurantId  string
	Cuisine       Cuisine
	CostBracket   int
	Rating        float32
	IsRecommended bool
	OnboardedTime time.Time
}

func GetRestaurantsForTheCuisineTypes(cuisines []Cuisine, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for _, cuisine := range cuisines {
		for _, restaurant := range restaurants {
			if restaurant.Cuisine == cuisine {
				filteredRestaurants = append(filteredRestaurants, restaurant)
			}
		}
	}
	return filteredRestaurants
}
func GetRestaurantsForTheCostBracket(costBrackets []int, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for _, costBracket := range costBrackets {
		for _, restaurant := range restaurants {
			if restaurant.CostBracket == costBracket {
				filteredRestaurants = append(filteredRestaurants, restaurant)
			}
		}
	}
	return filteredRestaurants
}
func GetFeaturedRestaurants(restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for _, restaurant := range restaurants {
		if restaurant.IsRecommended {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}
	return filteredRestaurants
}
func GetNewRestaurants(restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for _, restaurant := range restaurants {
		onBoardingTime := restaurant.OnboardedTime
		currentTime := time.Now()
		timeElapsed := currentTime.Sub(onBoardingTime)
		if (timeElapsed.Hours() / 24) <= 2 {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}
	return filteredRestaurants
}
func GetRestaurantsBelowGivenRating(rating float32, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for _, restaurant := range restaurants {
		if restaurant.Rating < rating {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}
	return filteredRestaurants
}
func GetRestaurantsAboveOrEqualGivenRating(rating float32, restaurants []Restaurant) []Restaurant {
	var filteredRestaurants []Restaurant
	for _, restaurant := range restaurants {
		if restaurant.Rating >= rating {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		}
	}
	return filteredRestaurants
}
