package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func PersonInit() {
	Friends := []Person{{"Ivan", 30}, {"Pavel", 28}}
	mapFriends := getPersonsInfo(Friends)

	fmt.Println(mapFriends)
}

func getPersonsInfo(personsSlice []Person) map[string]int {
	mapPersons := make(map[string]int)
	for _, person := range personsSlice {
		mapPersons[person.Name] = person.Age
	}
	return mapPersons
}
