package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

//All restaurants of secondary cuisine, primary cost bracket with rating >= 4.5

type SecondaryCuisinePrimaryCost struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitSecondaryCuisinePrimaryCost() SecondaryCuisinePrimaryCost {
	return SecondaryCuisinePrimaryCost{}
}

// GetRestaurants for PrimaryCuisinePrimaryCost fetches All restaurants of Primary cuisine, primary cost bracket with rating >= 4
func (p *SecondaryCuisinePrimaryCost) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	restaurantsWithPrimaryCuisine := domain.GetRestaurantsForTheCuisineTypes(userPreference.UserCuisinePreference.SecondaryCuisine, restaurants)
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
func (p *SecondaryCuisinePrimaryCost) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
