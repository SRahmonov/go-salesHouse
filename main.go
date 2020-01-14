package main

import "sort"

type house struct {
	id                int64
	title             string
	description       string
	image             []string
	top               bool
	price             int64
	city              string
	farFromCenterInKm int64
	district          string
}
func sortByPriceCheap(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}

func sortByPriceExpensive(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}

func sortByRemotenessFromCenterNear(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].farFromCenterInKm < result[j].farFromCenterInKm
	})
	return result
}
func sortByRemotenessFromCenterFar(houses []house) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].farFromCenterInKm > result[j].farFromCenterInKm
	})
	return result
}
func searchByMaxPrice(houses []house, max int64) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.price <= max {
			result = append(result, house)
		}
	}
	return result
}
func searchByMinAndMaxPrice(houses []house, min int64, max int64) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.price >= min && house.price <= max {
			result = append(result, house)
		}
	}
	return result
}
func searchByDistrict(houses []house, district string) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.district == district {
			result = append(result, house)
		}
	}
	return result
}
func searchByDistricts(houses []house, districts []string) []house {
	result := make([]house, 0)
	for _, house := range houses {
		for _, district := range districts {
			if house.district == district {
				result = append(result, house)
				continue
			}
		}
	}
	return result
}
func main() {
}
