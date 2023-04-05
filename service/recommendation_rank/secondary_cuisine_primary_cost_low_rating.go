package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

const MAX_ALLOWED_RATING_SECONDARY_CUISINE = 4.5

type SecondaryCuisineSecondaryCostLowRating struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitSecondaryCuisineSecondaryCostLowRating() SecondaryCuisineSecondaryCostLowRating {
	return SecondaryCuisineSecondaryCostLowRating{}
}

// GetRestaurants for PrimaryCuisineSecondaryCostLowRating fetches all restaurants of secondary cuisine, primary cost bracket with rating < 4.5
func (p *SecondaryCuisineSecondaryCostLowRating) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	restaurantsWithSecondaryCuisine := domain.GetRestaurantsForTheCuisineTypes(userPreference.UserCuisinePreference.SecondaryCuisine, restaurants)
	restaurantsInCostBracket := domain.GetRestaurantsForTheCostBracket([]int{userPreference.UserCostPreference.PrimaryCostbracket}, restaurantsWithSecondaryCuisine)
	restaurantsWithRating := domain.GetRestaurantsBelowGivenRating(MAX_ALLOWED_RATING_SECONDARY_CUISINE, restaurantsInCostBracket)
	for _, restaurant := range restaurantsWithRating {
		recommendedRestaurants = append(recommendedRestaurants, restaurant.RestaurantId)
	}
	if p.nextRecommendedRestaurants != nil {
		lowRankRecommendedRestaurants, err := p.nextRecommendedRestaurants.GetRestaurants(userPreference, restaurants)
		if err != nil {

		}
		recommendedRestaurants = utils.ConcatenateRecommendedRestaurants(recommendedRestaurants, lowRankRecommendedRestaurants)
	}
	return recommendedRestaurants, nil
}
func (p *SecondaryCuisineSecondaryCostLowRating) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
