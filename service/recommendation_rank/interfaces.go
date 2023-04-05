package recommendation_rank

import "RecommendationEngine/domain"

type RecommendationRank interface {
	GetRestaurants(preference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error)
	SetNextRank(rank RecommendationRank) error
}
