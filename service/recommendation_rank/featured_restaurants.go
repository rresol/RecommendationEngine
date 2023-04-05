package recommendation_rank

import (
	"RecommendationEngine/domain"
	"RecommendationEngine/utils"
)

type FeaturedRestaurants struct {
	nextRecommendedRestaurants RecommendationRank
}

func InitFeaturedRestaurants() FeaturedRestaurants {
	return FeaturedRestaurants{}
}

// GetRestaurants for FeaturedRestaurants fetches Featured restaurants of primary cuisine and primary cost bracket.
//
//	If none, then all featured restaurants of primary cuisine, secondary cost and secondary cuisine, primary cost
func (f *FeaturedRestaurants) GetRestaurants(userPreference domain.UserPreference, restaurants []domain.Restaurant) ([]string, error) {
	var recommendedFeaturedRestaurants []string
	featuredRestaurants := domain.GetFeaturedRestaurants(restaurants)
	featuredRestaurantsPrimaryCuisine := domain.GetRestaurantsForTheCuisineTypes([]domain.Cuisine{userPreference.UserCuisinePreference.PrimaryCuisine}, featuredRestaurants)
	featuredRestaurantsPrimaryCuisineCostBracket := domain.GetRestaurantsForTheCostBracket([]int{userPreference.UserCostPreference.PrimaryCostbracket}, featuredRestaurantsPrimaryCuisine)
	if len(featuredRestaurantsPrimaryCuisineCostBracket) > 0 {
		for _, restaurant := range featuredRestaurantsPrimaryCuisineCostBracket {
			recommendedFeaturedRestaurants = append(recommendedFeaturedRestaurants, restaurant.RestaurantId)
		}

	} else {
		featuredRestaurantsSecondaryCost := domain.GetRestaurantsForTheCostBracket(userPreference.UserCostPreference.SecondaryCostbracket, featuredRestaurantsPrimaryCuisine)
		featuredRestaurantsSecondaryCuisine := domain.GetRestaurantsForTheCuisineTypes(userPreference.UserCuisinePreference.SecondaryCuisine, featuredRestaurants)
		featuredRestaurantsSecondaryCuisineSecondaryCost := domain.GetRestaurantsForTheCostBracket([]int{userPreference.UserCostPreference.PrimaryCostbracket}, featuredRestaurantsSecondaryCuisine)
		featuredRestaurantsSecondaryCost = append(featuredRestaurantsSecondaryCost, featuredRestaurantsSecondaryCuisineSecondaryCost...)
		for _, restaurant := range featuredRestaurantsSecondaryCost {
			recommendedFeaturedRestaurants = append(recommendedFeaturedRestaurants, restaurant.RestaurantId)
		}
	}
	if f.nextRecommendedRestaurants != nil {
		lowRankRecommendedRestaurants, err := f.nextRecommendedRestaurants.GetRestaurants(userPreference, restaurants)
		if err != nil {

		}
		recommendedFeaturedRestaurants = utils.ConcatenateRecommendedRestaurants(recommendedFeaturedRestaurants, lowRankRecommendedRestaurants)
	}

	return recommendedFeaturedRestaurants, nil
}
func (f *FeaturedRestaurants) SetNextRank(rank RecommendationRank) error {
	f.nextRecommendedRestaurants = rank
	return nil
}
