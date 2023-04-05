package main

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/service/recommendation"
	"RecommendationEngine/service/recommendation_rank"
	"fmt"
	"time"
)

func main() {

	// initialise recommendation ranks
	featuredRestaurants := recommendation_rank.InitFeaturedRestaurants()
	primaryCuisineRestaurants := recommendation_rank.InitPrimaryCuisinePrimaryCost()
	primaryCuisineLowRating := recommendation_rank.InitPrimaryCuisinePrimaryCostLowRating()
	primaryCuisineSecondaryCost := recommendation_rank.InitPrimaryCuisineSecondaryCost()
	primaryCuisineSecondaryCostLowRating := recommendation_rank.InitPrimaryCuisineSecondaryCostLowRating()
	recentlyOnboardedRestaurants := recommendation_rank.InitNewOnboardedRestaurant()
	secondaryCuisinePrimaryCost := recommendation_rank.InitSecondaryCuisinePrimaryCost()
	secondaryCuisineSecondaryCost := recommendation_rank.InitSecondaryCuisineSecondaryCostLowRating()
	lowesRecommendedRestaurants := recommendation_rank.InitLowestRankedRestaurants()

	// define sort order for the recommended results
	featuredRestaurants.SetNextRank(&primaryCuisineRestaurants)
	primaryCuisineRestaurants.SetNextRank(&primaryCuisineSecondaryCost)
	primaryCuisineSecondaryCost.SetNextRank(&secondaryCuisinePrimaryCost)
	secondaryCuisinePrimaryCost.SetNextRank(&recentlyOnboardedRestaurants)
	recentlyOnboardedRestaurants.SetNextRank(&primaryCuisineLowRating)
	primaryCuisineLowRating.SetNextRank(&primaryCuisineSecondaryCostLowRating)
	primaryCuisineSecondaryCostLowRating.SetNextRank(&secondaryCuisineSecondaryCost)
	primaryCuisineSecondaryCostLowRating.SetNextRank(&lowesRecommendedRestaurants)

	// primary cuisine south Indian
	// secondary cuisine chinese
	// primary cost 5
	// seconary cost 4

	// create restaurants
	var availableRestaurants []domain.Restaurant
	recommendedRestaurantPrimaryCuisineHighCostHighRated := domain.Restaurant{
		RestaurantId:  "recommended_south_indian_restaurant_high_cost_high_rated_1",
		Cuisine:       domain.SouthIndian,
		CostBracket:   5,
		Rating:        5.0,
		IsRecommended: true,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, recommendedRestaurantPrimaryCuisineHighCostHighRated)
	// create restaurants
	southIndianRestaurantHighCostHighRated := domain.Restaurant{
		RestaurantId:  "south_indian_restaurant_high_cost_high_rated_1",
		Cuisine:       domain.SouthIndian,
		CostBracket:   5,
		Rating:        5.0,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, southIndianRestaurantHighCostHighRated)
	// create restaurants
	southIndianRestaurantMedCostHighRated := domain.Restaurant{
		RestaurantId:  "south_indian_restaurant_medium_cost_high_rated_1",
		Cuisine:       domain.SouthIndian,
		CostBracket:   4,
		Rating:        4.5,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, southIndianRestaurantMedCostHighRated)
	// create restaurants
	southIndianRestaurantHighCostLowRated := domain.Restaurant{
		RestaurantId:  "south_indian_restaurant_high_cost_low_rated_1",
		Cuisine:       domain.SouthIndian,
		CostBracket:   5,
		Rating:        3.0,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, southIndianRestaurantHighCostLowRated)

	// create restaurants
	chineseRestaurantMediumCost := domain.Restaurant{
		RestaurantId:  "chinese_restaurant_medium_cost_high_rated_1",
		Cuisine:       domain.Chinese,
		CostBracket:   5,
		Rating:        4.5,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, chineseRestaurantMediumCost)
	// create restaurants
	chineseRestaurantMediumCost2 := domain.Restaurant{
		RestaurantId:  "chinese_restaurant_medium_cost_high_rated_2",
		Cuisine:       domain.Chinese,
		CostBracket:   5,
		Rating:        4.5,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, chineseRestaurantMediumCost2)
	// create restaurants
	chineseRestaurantHighCostLowRating := domain.Restaurant{
		RestaurantId:  "chinese_restaurant_high_cost_low_rated_1",
		Cuisine:       domain.Chinese,
		CostBracket:   5,
		Rating:        2.0,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, chineseRestaurantHighCostLowRating)
	// create restaurants
	recentlyOnboardedRestaurants5 := domain.Restaurant{
		RestaurantId:  "recent_chinese_restaurant_high_cost_low_rated_1",
		Cuisine:       domain.Chinese,
		CostBracket:   5,
		Rating:        2.0,
		IsRecommended: false,
		OnboardedTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, recentlyOnboardedRestaurants5)
	// create restaurants
	recentlyOnboardedRestaurants1 := domain.Restaurant{
		RestaurantId:  "recent_north_indian_restaurant_high_cost_low_rated_1",
		Cuisine:       domain.NorthIndian,
		CostBracket:   5,
		Rating:        3.0,
		IsRecommended: false,
		OnboardedTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, recentlyOnboardedRestaurants1)
	// create restaurants
	recentlyOnboardedRestaurants2 := domain.Restaurant{
		RestaurantId:  "recent_south_indian_restaurant_high_cost_medium_rated_1",
		Cuisine:       domain.SouthIndian,
		CostBracket:   5,
		Rating:        1.0,
		IsRecommended: false,
		OnboardedTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, time.Now().Hour(), time.Now().Minute(), 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, recentlyOnboardedRestaurants2)

	lowPriorityRestaurants := domain.Restaurant{
		RestaurantId:  "low_priority_north_indian_restaurant",
		Cuisine:       domain.NorthIndian,
		CostBracket:   2,
		Rating:        2.0,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, lowPriorityRestaurants)

	lowPriorityRestaurants2 := domain.Restaurant{
		RestaurantId:  "low_priority_north_indian_restaurant_2",
		Cuisine:       domain.NorthIndian,
		CostBracket:   2,
		Rating:        2.5,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, lowPriorityRestaurants2)
	lowPriorityRestaurants5 := domain.Restaurant{
		RestaurantId:  "low_priority_north_indian_restaurant_3",
		Cuisine:       domain.NorthIndian,
		CostBracket:   2,
		Rating:        3.0,
		IsRecommended: false,
		OnboardedTime: time.Date(2022, 10, 2, 0, 0, 0, 0, time.UTC),
	}
	availableRestaurants = append(availableRestaurants, lowPriorityRestaurants5)

	trackedCuisinePrimary := domain.CuisineTracking{
		CuisineType: domain.SouthIndian,
		NoOfOrders:  100,
	}
	trackedCuisineSecondary := domain.CuisineTracking{
		CuisineType: domain.Chinese,
		NoOfOrders:  10,
	}
	trackedCostPrimary := domain.CostTracking{
		CostTrackingType: 5,
		NoOfOrders:       100,
	}
	trackedCostSecondary := domain.CostTracking{
		CostTrackingType: 4,
		NoOfOrders:       10,
	}
	user := domain.User{
		TrackedCuisines: []domain.CuisineTracking{trackedCuisineSecondary, trackedCuisinePrimary},
		TrackedCosts:    []domain.CostTracking{trackedCostPrimary, trackedCostSecondary},
	}

	// initialise recommendation engine
	recommendationEngine := recommendation.InitRecommendationService(&featuredRestaurants)

	restaurants, err := recommendationEngine.GetRestaurantRecommendations(user, availableRestaurants)
	if err != nil {
		fmt.Println("error while fetchin recommendation due to ", err)
	}
	for _, restaurant := range restaurants {
		fmt.Println("recommended restaurant: \n", restaurant)
	}

}
