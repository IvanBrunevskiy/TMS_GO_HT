package main

import "fmt"

type Product struct {
	Name     string
	Category string
	Price    float64
}

type ProductWithoutCategory struct {
	Name  string
	Price float64
}

func ProductsInit() {
	Products := []Product{{Name: "Laptop", Category: "Electronics", Price: 1000.5}, {Name: "Shirt", Category: "Clothes", Price: 25.5}}
	mapProducts := getProductsInfo(Products)

	fmt.Println(mapProducts)
}

func getProductsInfo(productsSlice []Product) map[string]ProductWithoutCategory {
	mapProducts := make(map[string]ProductWithoutCategory)
	for _, product := range productsSlice {
		mapProducts[product.Category] = ProductWithoutCategory{Name: product.Name, Price: product.Price}
	}
	return mapProducts
}
