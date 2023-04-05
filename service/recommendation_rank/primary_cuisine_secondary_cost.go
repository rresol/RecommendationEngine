package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

const MINIMUM_RESTAURANT_RATING_SECONDARY_COST = 4.5

type PrimaryCuisineSecondaryCost struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitPrimaryCuisineSecondaryCost() PrimaryCuisineSecondaryCost {
	return PrimaryCuisineSecondaryCost{}
}

// GetRestaurants for PrimaryCuisinePrimaryCost fetches All restaurants of Primary cuisine, primary cost bracket with rating >= 4
func (p *PrimaryCuisineSecondaryCost) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	restaurantsWithPrimaryCuisine := domain.GetRestaurantsForTheCuisineTypes([]domain.Cuisine{userPreference.UserCuisinePreference.PrimaryCuisine}, restaurants)
	restaurantsInCostBracket := domain.GetRestaurantsForTheCostBracket(userPreference.UserCostPreference.SecondaryCostbracket, restaurantsWithPrimaryCuisine)
	restaurantsWithRating := domain.GetRestaurantsAboveOrEqualGivenRating(MINIMUM_RESTAURANT_RATING_SECONDARY_COST, restaurantsInCostBracket)
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
func (p *PrimaryCuisineSecondaryCost) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
