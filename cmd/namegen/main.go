package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ironarachne/namegen"
	"github.com/ironarachne/random"
)

func main() {
	numberOfNames := flag.Int("n", 1, "Number of names to generate")
	gender := flag.String("g", "both", "Gender (male, female, or both)")
	origin := flag.String("o", "english", "Origin of names (e.g., english, spanish, etc.)")
	randomSeed := flag.String("s", "none", "Optional random generator seed (alphanumeric)")

	mode := flag.String("m", "full", "Mode of generation")
	var list = flag.Bool("l", false, "Print available name lists")

	flag.Parse()

	if *list == true {
		nameLists := []string{
			"anglosaxon",
			"dutch",
			"elf",
			"english",
			"fantasy",
			"german",
			"greek",
			"indonesian",
			"japanese",
			"korean",
			"russian",
			"spanish",
		}
		fmt.Printf("Available name lists: \n%s\n\n", strings.Join(nameLists, "\n"))
	}

	if *randomSeed == "none" {
		rand.Seed(time.Now().UnixNano())
	} else {
		random.SeedFromString(*randomSeed)
	}

	generator := namegen.NameGeneratorFromType(*origin)

	output := ""
	for i := 0; i < *numberOfNames; i++ {

		switch *mode {
		case "full":
			output = generator.CompleteName(*gender)
		case "firstname":
			output = generator.FirstName(*gender)
		case "lastname":
			output = generator.LastName()
		default:
			output = fmt.Sprintf("Unsupported generation mode %s, supported modes: full, firstname, lastname", *mode)
		}

		fmt.Println(output)
	}
}
