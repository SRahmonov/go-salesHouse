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

func sortByAscAndDesc(houses []house) (resultAsc, resultDesc []house) {
	resultAsc = make([]house, len(houses))
	copy(resultAsc, houses)
	resultDesc = make([]house, len(houses))
	copy(resultDesc, houses)
	sort.Slice(resultAsc, func(i, j int) bool {
		return resultAsc[i].price < resultAsc[j].price
	})
	sort.Slice(resultDesc, func(i, j int) bool {
		return resultDesc[i].price > resultDesc[j].price
	})
	return
}

func sortByRemotenessFromCenterNearAndFar(houses []house) (resultForNear, resultForFar []house) {
	resultForNear = make([]house, len(houses))
	copy(resultForNear, houses)
	resultForFar = make([]house, len(houses))
	copy(resultForFar, houses)
	sort.Slice(resultForNear, func(i, j int) bool {
		return resultForNear[i].farFromCenterInKm < resultForNear[j].farFromCenterInKm
	})
	sort.Slice(resultForFar, func(i, j int) bool {
		return resultForFar[i].farFromCenterInKm > resultForFar[j].farFromCenterInKm
	})
	return
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
