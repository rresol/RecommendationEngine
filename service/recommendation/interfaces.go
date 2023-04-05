package recommendation

import "RecommendationEngine/domain"

type RecommendationService interface {
	GetRestaurantRecommendations(user domain.User, restaurants []domain.Restaurant) ([]string, error)
}
