package main

import "fmt"

var houses = []house{
	{
		id:                1,
		title:             "2-комн. квартира, 1 этаж, 34 м²",
		description:       "Кирпичный дом. Имеется гараж",
		image:             []string{},
		top:               false,
		price:             90_000,
		city:              "Душанбе",
		farFromCenterInKm: 5,
		district:          "Рудаки",
	},
	{
		id:                2,
		title:             "3-комн. квартира, 2 этаж, 42 м²",
		description:       "Кирпичный дом. Имеется огород.",
		image:             []string{},
		top:               false,
		price:             150_000,
		city:              "Душанбе",
		farFromCenterInKm: 1,
		district:          "Рудаки",
	},
	{
		id:                3,
		title:             "2-комн. квартира, 5 этаж, 25 м²",
		description:       "Кирпичный дом. Имеется огород.",
		image:             []string{},
		top:               false,
		price:             120_000,
		city:              "Душанбе",
		farFromCenterInKm: 0,
		district:          "Фирдавси",
	},
}
func ExampleSortByPriceCheap() {
	ascSorted := sortByPriceCheap(houses)
	fmt.Println(ascSorted)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки} {3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси} {2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки}]
}

func ExampleSortByPriceExpensive() {
	descSorted := sortByPriceExpensive(houses)
	fmt.Println(descSorted)
	// Output: [{2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки} {3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси} {1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки}]
}

func ExampleSearchByMaxPrice_NoResults() {
	result := searchByMaxPrice(houses, 80_000)
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByMaxPrice_OneResult() {
	result := searchByMaxPrice(houses, 100_500)
	fmt.Println(result)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки}]
}

func ExampleSearchByMaxPrice_TwoAndMoreResults() {
	result := searchByMaxPrice(houses, 170_000)
	fmt.Println(result)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки} {2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки} {3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси}]
}

func ExampleSearchByMinAndMaxPrice_NoResults() {
	result := searchByMinAndMaxPrice(houses, 10_000, 10_500)
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByMinAndMaxPrice_OneResult() {
	result := searchByMinAndMaxPrice(houses, 50_000, 100_500)
	fmt.Println(result)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки}]
}

func ExampleSearchByMinAndMaxPrice_TwoOrMoreResults() {
	result := searchByMinAndMaxPrice(houses, 90_000, 120_500)
	fmt.Println(result)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки} {3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси}]
}

func ExampleSearchByDistrict_NoResults() {
	result := searchByDistrict(houses, "Айни")
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByDistrict_OneResult() {
	result := searchByDistrict(houses, "Фирдавси")
	fmt.Println(result)
	// Output: [{3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси}]
}

func ExampleSearchByDistrict_TwoOrMoreResults() {
	result := searchByDistrict(houses, "Рудаки")
	fmt.Println(result)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки} {2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки}]
}

func ExampleSearchByDistricts_NoResults() {
	result := searchByDistricts(houses, []string{"Айни", "Шохмансур"})
	fmt.Println(result)
	// Output: []
}

func ExampleSearchByDistricts_OneResult() {
	result := searchByDistricts(houses, []string{"Фирдавси", "Шохмансур"})
	fmt.Println(result)
	// Output: [{3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси}]
}

func ExampleSearchByDistricts_TwoOrMoreResults() {
	result := searchByDistricts(houses, []string{"Фирдавси", "Рудаки"})
	fmt.Println(result)
	// Output: [{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки} {2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки} {3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси}]
}

func ExampleSortByRemotenessFromCenterNear() {
	ascSorted := sortByRemotenessFromCenterNear(houses)
	fmt.Println(ascSorted)
	// Output: [{3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси} {2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки}{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж. [] false 90000 Душанбе 5 Рудаки}]
}

func ExampleSortByRemotenessFromCenterFar() {
	descSorted := sortByRemotenessFromCenterFar(houses)
	fmt.Println(descSorted)
	// Output:[{1 2-комн. квартира, 1 этаж, 34 м² Кирпичный дом. Имеется гараж [] false 90000 Душанбе 5 Рудаки} {2 3-комн. квартира, 2 этаж, 42 м² Кирпичный дом. Имеется огород. [] false 150000 Душанбе 1 Рудаки} {3 2-комн. квартира, 5 этаж, 25 м² Кирпичный дом. Имеется огород. [] false 120000 Душанбе 0 Фирдавси}]
}