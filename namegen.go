package namegen

import (
	"github.com/ironarachne/random"
)

// NameGenerator is a set of names to use
type NameGenerator struct {
	MaleFirstNames   []string
	FemaleFirstNames []string
	LastNames        []string
}

// NameGeneratorFromType sets up types of names
func NameGeneratorFromType(origin, gender string) NameGenerator {
	nameGenerators := map[string]NameGenerator{
		"anglosaxon": {anglosaxonMaleFirstNames, anglosaxonFemaleFirstNames, anglosaxonLastNames},
		"dutch":      {dutchMaleFirstNames, dutchFemaleFirstNames, dutchLastNames},
		"dwarf": 	  {dwarfMaleFirstNames, dwarfFemaleFirstNames, getDwarfLastNames(gender)},
		"elf":        {elfMaleFirstNames, elfFemaleFirstNames, elfLastNames},
		"english":    {englishMaleFirstNames, englishFemaleFirstNames, englishLastNames},
		"fantasy":    {fantasyMaleFirstNames, fantasyFemaleFirstNames, fantasyLastNames},
		"german":     {germanMaleFirstNames, germanFemaleFirstNames, germanLastNames},
		"greek":      {greekMaleFirstNames, greekFemaleFirstNames, greekLastNames},
		"icelandic":  {getIcelandicFirstNames(), getIcelandicFirstNames(), getIcelandicLastNames(gender)},
		"indonesian": {indonesianMaleFirstNames, indonesianFemaleFirstNames, indonesianLastNames},
		"japanese":   {japaneseMaleFirstNames, japaneseFemaleFirstNames, japaneseLastNames},
		"korean":     {koreanMaleFirstNames, koreanFemaleFirstNames, koreanLastNames},
		"spanish":    {spanishMaleFirstNames, spanishFemaleFirstNames, spanishLastNames},
		"russian":    {russianMaleFirstNames, russianFemaleFirstNames, russianLastNames},
	}

	return nameGenerators[origin]
}

// LastName returns a last name
func (gen NameGenerator) LastName() string {
	return random.Item(gen.LastNames)
}

// FirstName returns a first name
func (gen NameGenerator) FirstName(gender string) string {
	firstNames := gen.MaleFirstNames
	if gender == "female" {
		firstNames = gen.FemaleFirstNames
	} else if gender == "both" {
		firstNames = append(firstNames, gen.FemaleFirstNames...)
	}

	return random.Item(firstNames)
}

// CompleteName returns a complete name
func (gen NameGenerator) CompleteName(gender string) string {
	return gen.FirstName(gender) + " " + gen.LastName()
}
