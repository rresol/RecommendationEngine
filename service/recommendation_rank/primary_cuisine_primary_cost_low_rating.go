package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

//All restaurants of Primary cuisine, primary cost bracket with rating < 4

const MAX_ALLOWED_RATING_PRIMARY_COST = 4.0

type PrimaryCuisinePrimaryCostLowRating struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitPrimaryCuisinePrimaryCostLowRating() PrimaryCuisinePrimaryCostLowRating {
	return PrimaryCuisinePrimaryCostLowRating{}
}

// GetRestaurants for PrimaryCuisinePrimaryCost fetches All restaurants of Primary cuisine, primary cost bracket with rating >= 4
func (p *PrimaryCuisinePrimaryCostLowRating) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	restaurantsWithPrimaryCuisine := domain.GetRestaurantsForTheCuisineTypes([]domain.Cuisine{userPreference.UserCuisinePreference.PrimaryCuisine}, restaurants)
	restaurantsInCostBracket := domain.GetRestaurantsForTheCostBracket([]int{userPreference.UserCostPreference.PrimaryCostbracket}, restaurantsWithPrimaryCuisine)
	restaurantsWithRating := domain.GetRestaurantsBelowGivenRating(MAX_ALLOWED_RATING_PRIMARY_COST, restaurantsInCostBracket)
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
func (p *PrimaryCuisinePrimaryCostLowRating) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
