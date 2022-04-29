package utils

import (
	"mgidAssignment/pkg/models"
	"unicode"
)

type ByAge []models.Employee
type ByName []models.Employee
type BySalary []models.Employee

func (e ByAge) Len() int {
	return len(e)
}

func (e ByAge) Less(i, j int) bool {
	return e[i].YearOfBirth < e[j].YearOfBirth
}

func (e ByAge) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e ByName) Len() int {
	return len(e)
}

func (e ByName) Less(i, j int) bool {
	iRunes := []rune(e[i].Name)
	jRunes := []rune(e[j].Name)

	max := len(iRunes)
	if max > len(jRunes) {
		max = len(jRunes)
	}

	for idx := 0; idx < max; idx++ {
		ir := iRunes[idx]
		jr := jRunes[idx]

		lir := unicode.ToLower(ir)
		ljr := unicode.ToLower(jr)

		if lir != ljr {
			return lir < ljr
		}

		// the lowercase runes are the same, so compare the original
		if ir != jr {
			return ir < jr
		}
	}

	// If the strings are the same up to the length of the shortest string,
	// the shorter string comes first
	return len(iRunes) < len(jRunes)
}

func (e ByName) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e BySalary) Len() int {
	return len(e)
}

func (e BySalary) Less(i, j int) bool {
	return e[i].Salary > e[j].Salary
}

func (e BySalary) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
