package main

import "fmt"

type Student struct {
	Name        string
	Speciality  string
	Enrollment  int // Год поступления
	AverageMark float64
}

func StudentsInit() {
	StudentsInfoSlice := []Student{
		{Name: "Alex", Speciality: "Math", Enrollment: 2020, AverageMark: 4.5},
		{Name: "Alex", Speciality: "Math", Enrollment: 2020, AverageMark: 14},
		{Name: "Bob", Speciality: "Math", Enrollment: 2021, AverageMark: 4.0},
		{Name: "Charlie", Speciality: "Physics", Enrollment: 2020, AverageMark: 4.8},
		{Name: "Charlie2", Speciality: "Physics", Enrollment: 2020, AverageMark: 10},
	}

	info := getStudentsInfo(StudentsInfoSlice)

	fmt.Println(info)
}

func getStudentsInfo(studentsSlice []Student) map[string]map[int]float64 {
	StudentsInfoMap := make(map[string]map[int]float64)

	for _, studentInfo := range studentsSlice {
		_, ok := StudentsInfoMap[studentInfo.Speciality]
		if !ok {
			StudentsInfoMap[studentInfo.Speciality] = make(map[int]float64)
			StudentsInfoMap[studentInfo.Speciality][studentInfo.Enrollment] = studentInfo.AverageMark
		} else {
			mark, ok := StudentsInfoMap[studentInfo.Speciality][studentInfo.Enrollment]
			if ok {
				value := 0.0
				value = (mark + studentInfo.AverageMark) / 2
				StudentsInfoMap[studentInfo.Speciality][studentInfo.Enrollment] = value
			} else {
				StudentsInfoMap[studentInfo.Speciality][studentInfo.Enrollment] = studentInfo.AverageMark
			}
		}
	}

	return StudentsInfoMap
}
