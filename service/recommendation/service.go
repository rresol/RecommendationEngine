package recommendation

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/service/recommendation_rank"
	"errors"
	"sort"
)

type RecommendationServiceImpl struct {
	recommendedRankedRestaurants recommendation_rank.RecommendationRank
}

func InitRecommendationService(recommendedRankedRestaurants recommendation_rank.RecommendationRank) RecommendationServiceImpl {
	return RecommendationServiceImpl{
		recommendedRankedRestaurants: recommendedRankedRestaurants,
	}
}

func (r *RecommendationServiceImpl) computeUserPreferences(user domain.User) (domain.UserPreference, error) {
	if len(user.TrackedCuisines) == 0 {
		return domain.UserPreference{}, errors.New("no cuisines tracked for user")
	}
	if len(user.TrackedCosts) == 0 {
		return domain.UserPreference{}, errors.New("no costs tracked for user")
	}

	costBrackets := user.TrackedCosts
	sort.Slice(costBrackets[:], func(i, j int) bool {
		return costBrackets[i].NoOfOrders > costBrackets[j].NoOfOrders
	})

	cusines := user.TrackedCuisines
	sort.Slice(cusines, func(i, j int) bool {
		return cusines[i].NoOfOrders > cusines[j].NoOfOrders
	})

	var secondaryCuisines []domain.Cuisine
	var secondaryCostBracket []int

	primaryCuisine := cusines[0].CuisineType
	primaryCostBracket := costBrackets[0].CostTrackingType

	if len(cusines) == 2 {
		secondaryCuisines = append(secondaryCuisines, cusines[1].CuisineType)
	}
	if len(cusines) > 2 {
		secondaryCuisines = append(secondaryCuisines, cusines[2].CuisineType)
	}
	if len(costBrackets) == 2 {
		secondaryCostBracket = append(secondaryCostBracket, costBrackets[1].CostTrackingType)
	}
	if len(costBrackets) > 2 {
		secondaryCostBracket = append(secondaryCostBracket, costBrackets[2].CostTrackingType)
	}
	userCuisinePreference := domain.UserCuisinePreference{
		PrimaryCuisine:   primaryCuisine,
		SecondaryCuisine: secondaryCuisines,
	}
	userCostPreference := domain.UserCostPreference{
		PrimaryCostbracket:   primaryCostBracket,
		SecondaryCostbracket: secondaryCostBracket,
	}
	userPreference := domain.UserPreference{
		UserCuisinePreference: userCuisinePreference,
		UserCostPreference:    userCostPreference,
	}
	return userPreference, nil
}
func (r *RecommendationServiceImpl) GetRestaurantRecommendations(user domain.User, restaurants []domain.Restaurant) ([]string, error) {
	userPreference, err := r.computeUserPreferences(user)
	if err != nil {
		return nil, err
	}

	recommendedRestaurants, gErr := r.recommendedRankedRestaurants.GetRestaurants(userPreference, restaurants)
	if gErr != nil {
		return nil, gErr
	}
	return recommendedRestaurants, nil
}
