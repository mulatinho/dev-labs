package utils

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type NAME_TYPE int

const (
	NAME_TYPE_MAN NAME_TYPE = iota
	NAME_TYPE_WOMAN
	NAME_TYPE_TOTAL
)

func GetPlayerID() int {
	rand.Seed(time.Now().UnixMicro())

	return rand.Int() % 20
}

func GetManName() string {
	NamesList := []string{
		"John", "Jack", "Lemmy", "Binni", "Pierette",
		"Layne", "Donny", "Murial", "Hedwig", "Brenn",
		"Dion", "Kirby", "Effie", "Maddy", "Bob",
		"Reggie", "Albert", "Joseph", "Lane", "Joey", "Verne",
	}

	return NamesList[rand.Int()%len(NamesList)]
}

func GetWomanName() string {
	NamesList := []string{
		"Cathrin", "Zilvia", "Estele", "Norina", "Donny",
		"Kaitlynn", "Silvana", "Milka", "Evonne", "Jacquelynn",
		"Sofia", "Ariadne", "Laurice", "Millicent", "Alexa",
		"Leandra", "Simone", "Lily", "Marya", "Shayla",
	}

	return NamesList[rand.Int()%len(NamesList)]
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

	return NamesList[rand.Int()%len(NamesList)]
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

type LogType int

const (
	LOG_TYPE_NORMAL LogType = iota
	LOG_TYPE_DEBUG
	LOG_TYPE_PANIC
)

func LogMessage(logType LogType, message any) {
	_, fileName, lineNumber, _ := runtime.Caller(1)
	displayLogEnabled := bool(os.Getenv("DEBUG") != "")

	switch logType {
	case LOG_TYPE_PANIC:
		fmt.Printf(":. \033[1m PANIC %s:%d\033[m %#v\n", filepath.Base(fileName), lineNumber, message)
		os.Exit(1)
	case LOG_TYPE_DEBUG:
		if displayLogEnabled {
			fmt.Printf(":. \033[1m DEBUG %s:%d\033[m %#v\n", filepath.Base(fileName), lineNumber, message)
		}
	default:
		if displayLogEnabled {
			fmt.Printf(":. \033[1mNORMAL\033[m %v\n", message)
		}
	}
}
