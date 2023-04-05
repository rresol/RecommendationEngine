package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

const MAX_ALLOWED_RATING_SECONDARY_COST = 4.5

type PrimaryCuisineSecondaryCostLowRating struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitPrimaryCuisineSecondaryCostLowRating() PrimaryCuisineSecondaryCostLowRating {
	return PrimaryCuisineSecondaryCostLowRating{}
}

// GetRestaurants for PrimaryCuisineSecondaryCostLowRating fetches All restaurants of Primary cuisine, secondary cost bracket with rating < 4.5
func (p *PrimaryCuisineSecondaryCostLowRating) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	restaurantsWithPrimaryCuisine := domain.GetRestaurantsForTheCuisineTypes([]domain.Cuisine{userPreference.UserCuisinePreference.PrimaryCuisine}, restaurants)
	restaurantsInCostBracket := domain.GetRestaurantsForTheCostBracket(userPreference.UserCostPreference.SecondaryCostbracket, restaurantsWithPrimaryCuisine)
	restaurantsWithRating := domain.GetRestaurantsBelowGivenRating(MAX_ALLOWED_RATING_SECONDARY_COST, restaurantsInCostBracket)
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
func (p *PrimaryCuisineSecondaryCostLowRating) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
