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

func sortBy(houses []house, less func(a, b house) bool) []house {
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}
func sortByPriceAsc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price < b.price
	})

}
func sortByPriceDesc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.price > b.price
	})
}
func farFromCenterAsc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.farFromCenterInKm > b.farFromCenterInKm

	})
}
func farFromCenterDesc(houses []house) []house {
	return sortBy(houses, func(a, b house) bool {
		return a.farFromCenterInKm < b.farFromCenterInKm

	})
}
func searchBy(houses []house, less func(a house) bool) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if less(house) {
			result = append(result, house)
		}
	}
	return result
}
func searchByMaxPrice(houses []house, MaxPrice int64) []house {
	return searchBy(houses, func(a house) bool {
		return a.price <= MaxPrice

	})
}
func searchByPrice(houses []house, startLimit int64, endLimit int64) []house {
	return searchBy(houses, func(a house) bool {
		return a.price >= startLimit && a.price <= endLimit
	})
}

func searchByDistricts(houses []house, districts string) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.district == districts {
			result = append(result, house)
		}
	}
	return result
}
func main() {
}
