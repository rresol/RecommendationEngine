package recommendation_rank

import (
	"RecommendationEngine/domain"
)

type LowestRankedRestaurants struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitLowestRankedRestaurants() LowestRankedRestaurants {
	return LowestRankedRestaurants{}
}

// GetRestaurants for LowestRankedRestaurants fetches all restaurants
func (p *LowestRankedRestaurants) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRestaurants []string
	for _, restaurant := range restaurants {
		recommendedRestaurants = append(recommendedRestaurants, restaurant.RestaurantId)
	}
	return recommendedRestaurants, nil
}
func (p *LowestRankedRestaurants) SetNextRank(rank RecommendationRank) error {
	p.nextRecommendedRestaurants = rank
	return nil
}
