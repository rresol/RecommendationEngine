package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

const MINIMUM_RESTAURANT_RATING = 4.5

type PrimaryCuisinePrimaryCost struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitPrimaryCuisinePrimaryCost() PrimaryCuisinePrimaryCost {
	return PrimaryCuisinePrimaryCost{}
}

// GetRestaurants for PrimaryCuisinePrimaryCost fetches All restaurants of Primary cuisine, primary cost bracket with rating >= 4
func (p *PrimaryCuisinePrimaryCost) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	restaurantsWithPrimaryCuisine := domain.GetRestaurantsForTheCuisineTypes([]domain.Cuisine{userPreference.UserCuisinePreference.PrimaryCuisine}, restaurants)
	restaurantsInCostBracket := domain.GetRestaurantsForTheCostBracket([]int{userPreference.UserCostPreference.PrimaryCostbracket}, restaurantsWithPrimaryCuisine)
	restaurantsWithRating := domain.GetRestaurantsAboveOrEqualGivenRating(MINIMUM_RESTAURANT_RATING, restaurantsInCostBracket)
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
func (p *PrimaryCuisinePrimaryCost) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
