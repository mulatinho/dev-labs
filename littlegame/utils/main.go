package utils

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"golang.org/x/exp/slices"
)

type (
	NAME_TYPE int
	LogType   int
)

const (
	NAME_TYPE_MAN NAME_TYPE = iota
	NAME_TYPE_WOMAN
	NAME_TYPE_TOTAL
)

const (
	LOG_TYPE_NORMAL LogType = iota
	LOG_TYPE_DEBUG
	LOG_TYPE_PANIC
)

func GetRandomNumber(n ...int) int {
	var result int
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	if len(n) == 0 {
		result = r.Int() % 10_000
	} else {
		result = r.Int() % n[0]
	}
	return result
}

func GetManName() string {
	NamesList := []string{
		"John", "Jack", "Lemmy", "Binni", "Pierette",
		"Layne", "Donny", "Murial", "Hedwig", "Brenn",
		"Dion", "Kirby", "Effie", "Maddy", "Bob",
		"Reggie", "Albert", "Joseph", "Lane", "Joey", "Verne",
	}

	index := GetRandomNumber(len(NamesList))
	return NamesList[index]
}

func GetWomanName() string {
	NamesList := []string{
		"Cathrin", "Zilvia", "Estele", "Norina", "Donny",
		"Kaitlynn", "Silvana", "Milka", "Evonne", "Jacquelynn",
		"Sofia", "Ariadne", "Laurice", "Millicent", "Alexa",
		"Leandra", "Simone", "Lily", "Marya", "Shayla",
	}

	index := GetRandomNumber(len(NamesList))
	return NamesList[index]
}

func GetSurname() string {
	NamesList := []string{
		"Valencia", "Manning", "Dunn", "Wilson", "Black", "Vega", "Walls",
		"Savage", "Stevenson", "Thornton", "Johnson", "Gordon", "Steele", "Singleton",
		"Dillon", "Kane", "Case", "Joyce", "Mayo", "William", "Hansen",
		"Parsons", "Robinson", "Townsend", "Cote", "Barlow", "Burt", "Moss",
		"Lloyd", "Santos", "Anderson", "Rocha", "Patton", "Jordan", "Serrano",
		"Jacobs", "Benton", "Ellison", "Short", "Raymond", "Schmidt", "Daniels",
		"Wong", "Woodard", "Miranda", "Johnston", "Aguirre", "Higgins", "Rivas",
		"Murray", "Reilly", "Aguilar", "Rodriguez", "Witt", "Talley", "Barry",
		"Bradford", "Copeland", "Newman", "Velazquez", "Cohen", "Mueller", "Herring",
		"Kidd", "Mccormick", "Jefferson", "Young", "Hendricks", "Ford", "Morrow",
		"Wiggins", "Mckinney", "Oliver", "Farrell", "Beasley", "Charles", "Barton",
		"Dunlap", "Merritt", "Henderson", "Ortega", "Montgomery", "Bryant", "Wilcox",
		"Davidson", "Dodson", "Wallace", "Cox", "Pickett", "Jennings", "Noble",
		"Walsh", "Warner", "House", "May", "Vasquez", "Ward", "Hopper", "Landry", "Combs",
	}

	index := GetRandomNumber(len(NamesList))
	return NamesList[index]
}

func GenerateName(name_type NAME_TYPE) string {
	var newName string

	switch name_type {
	case NAME_TYPE_MAN:
		newName = fmt.Sprintf("%s %s", GetManName(), GetSurname())
	case NAME_TYPE_WOMAN:
		newName = fmt.Sprintf("%s %s", GetWomanName(), GetSurname())
	default:
		newName = "ANONYMOUS"
	}

	return newName
}

func LogMessage(logType LogType, message ...any) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	displayLogEnabled := bool(os.Getenv("DEBUG") != "")

	switch logType {
	case LOG_TYPE_PANIC:
		fmt.Printf(":. \033[1m PANIC %s:%d\033[m ", filepath.Base(fileName), lineNumber)
		fmt.Printf("%#v\n", message...)
		os.Exit(1)
	case LOG_TYPE_DEBUG:
		if displayLogEnabled {
			fmt.Printf(":. \033[1m DEBUG %s:%d\033[m ", filepath.Base(fileName), lineNumber)
			fmt.Printf("%#v\n", message...)
			//fmt.Printf(":. \033[1m DEBUG %s:%d\033[m %#v\n", filepath.Base(fileName), lineNumber, message)
		}
	default:
		if displayLogEnabled {
			fmt.Printf(":. \033[1mNORMAL %s:%d\033[m ", filepath.Base(fileName), lineNumber)
			fmt.Printf("%#v\n", message...)
			//fmt.Printf(":. \033[1mNORMAL\033[m %v\n", message...)
		}
	}
}

func ApplyRandomIndexesToArray(totalLen int) []int {
	var arrayResult []int

	for index := 0; index < totalLen; index++ {
		for {
			chosen := GetRandomNumber(totalLen)
			if !slices.Contains(arrayResult, chosen) {
				arrayResult = append(arrayResult, chosen)
				break
			}
		}
	}

	return arrayResult
}

func ApplyRandomIndexesToArrayPreBuilt(totalLen int, prebuilt []int) []int {
	var arrayResult []int

	arrayResult = append(arrayResult, prebuilt...)

	for index := len(prebuilt); index < totalLen; index++ {
		for {
			chosen := GetRandomNumber(totalLen)
			if !slices.Contains(arrayResult, chosen) {
				arrayResult = append(arrayResult, chosen)
				break
			}
		}
	}

	return arrayResult
}
