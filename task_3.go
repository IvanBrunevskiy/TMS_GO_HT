package main

import "fmt"

type Citizen struct {
	Name string
	Age  int
}

func CitizenInit() {
	CityInfoMap := make(map[string][]Citizen)
	CityInfoMap["CityA"] = append(CityInfoMap["CityA"], Citizen{Name: "Alex", Age: 25})
	CityInfoMap["CityA"] = append(CityInfoMap["CityA"], Citizen{Name: "Bob", Age: 25})
	CityInfoMap["CityB"] = append(CityInfoMap["CityB"], Citizen{Name: "Charlie", Age: 30})
	info := getCityInfo(CityInfoMap)

	fmt.Println(info)
}

func getCityInfo(cityMap map[string][]Citizen) map[int]int {
	mapInfo := make(map[int]int)
	for _, nameAgeSlice := range cityMap {
		for _, nameAge := range nameAgeSlice {
			_, ok := mapInfo[nameAge.Age]
			if ok {
				mapInfo[nameAge.Age] = mapInfo[nameAge.Age] + 1
			} else {
				mapInfo[nameAge.Age] = 1
			}
		}
	}
	return mapInfo
}
