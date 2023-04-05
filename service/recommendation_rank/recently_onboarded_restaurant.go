package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
	"sort"
)

const MAX_RECENT_RESTAURANTS_ALLOWED_FOR_RECOMMENDATION = 4

type NewOnboardedRestaurant struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitNewOnboardedRestaurant() NewOnboardedRestaurant {
	return NewOnboardedRestaurant{}
}

// GetRestaurants for NewOnboardedRestaurant fetches top 4 newly created restaurants by rating
func (n *NewOnboardedRestaurant) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedRecentOnboardedRestaurants []string
	recentOnboardedRestaurants := domain.GetNewRestaurants(restaurants)
	sort.Slice(recentOnboardedRestaurants, func(i, j int) bool {
		return recentOnboardedRestaurants[i].Rating > recentOnboardedRestaurants[j].Rating
	})
	for _, restaurant := range recentOnboardedRestaurants {
		recommendedRecentOnboardedRestaurants = append(recommendedRecentOnboardedRestaurants, restaurant.RestaurantId)
		if len(recommendedRecentOnboardedRestaurants) == MAX_RECENT_RESTAURANTS_ALLOWED_FOR_RECOMMENDATION {
			break
		}
	}
	if n.nextRecommendedRestaurants != nil {
		lowRankRecommendedRestaurants, err := n.nextRecommendedRestaurants.GetRestaurants(userPreference, restaurants)
		if err != nil {

		}
		recommendedRecentOnboardedRestaurants = utils.ConcatenateRecommendedRestaurants(recommendedRecentOnboardedRestaurants, lowRankRecommendedRestaurants)
	}
	return recommendedRecentOnboardedRestaurants, nil
}
func (n *NewOnboardedRestaurant) SetNextRank(rank RecommendationRank) error {
	n.nextRecommendedRestaurants = rank
	return nil
}
