package stats

import (
	"log"
	database "mysql"
)

func GetTotalUsers() int {
	count, err := database.CountTableRows("users")
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func GetTotalOrganizations() int {
	count, err := database.CountTableRows("organizations")
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func GetTotalProjects() int {
	count, err := database.CountTableRows("projects")
	if err != nil {
		log.Fatal(err)
	}
	return count
}
