package main

import (
	"regexp"
	"sort"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

// returns part1 and part2
func run(input string) (interface{}, interface{}) {
	recipes := parse(input)

	allergens := make(map[string]*Allergen)
	words := make(map[string]int)
	for _, recipe := range recipes {
		for _, word := range recipe.Ingredients {
			words[word]++
		}
		for _, allergen := range recipe.Allergens {
			if _, exist := allergens[allergen]; !exist {
				allergens[allergen] = &Allergen{
					Name:  allergen,
					CanBe: make(map[string]struct{}),
				}
				for _, ingredient := range recipe.Ingredients {
					allergens[allergen].CanBe[ingredient] = struct{}{}
				}
				continue
			}
			intersections := make(map[string]struct{})
			for _, ingredient1 := range recipe.Ingredients {
				for Ingredient2 := range allergens[allergen].CanBe {
					if ingredient1 == Ingredient2 {
						intersections[Ingredient2] = struct{}{}
					}
				}
			}
			allergens[allergen].CanBe = intersections
		}
	}

	sortedAllergens := []*Allergen(nil)
	allSet := false
	for !allSet {
		allSet = true
		for _, allergen := range allergens {
			switch len(allergen.CanBe) {
			case 1:
				for i := range allergen.CanBe {
					allergen.Translation = i
				}
				for _, allergen2 := range allergens {
					delete(allergen2.CanBe, allergen.Translation)
				}
				sortedAllergens = append(sortedAllergens, allergen)
			case 0:
			default:
				allSet = false
			}
		}
	}

	part1 := 0
	for word, count := range words {
		isAllergen := false
		for _, allergen := range allergens {
			if allergen.Translation == word {
				isAllergen = true
				break
			}
		}
		if !isAllergen {
			part1 += count
		}
	}

	sort.Slice(sortedAllergens, func(i, j int) bool {
		return sortedAllergens[i].Name < sortedAllergens[j].Name
	})
	canonicalDangerousList := []string(nil)
	for _, allergen := range sortedAllergens {
		canonicalDangerousList = append(canonicalDangerousList, allergen.Translation)
	}

	return part1, strings.Join(canonicalDangerousList, ",")
}

type Allergen struct {
	Name        string
	Translation string
	CanBe       map[string]struct{}
}

type Recipe struct {
	Ingredients []string
	Allergens   []string
}

func parse(s string) []*Recipe {
	recipes := []*Recipe(nil)
	for _, matches := range regexp.MustCompile(`(.*) \(contains (.*)\)`).FindAllStringSubmatch(s, -1) {
		recipes = append(recipes, &Recipe{
			Ingredients: strings.Split(matches[1], " "),
			Allergens:   strings.Split(matches[2], ", "),
		})
	}
	return recipes
}

func main() {
	execute.Run(run, tests, puzzle, true)
}
