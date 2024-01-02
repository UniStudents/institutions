package main

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var importedInstitutions map[string]*DatabaseInstitution

func main() {
	importedInstitutions = make(map[string]*DatabaseInstitution)

	webInstitutions := GetInstitutionsFromWeb()
	databaseInstitutions := GetInstitutionsFromDatabase()

	for index := range webInstitutions {
		findAndProcess(index, webInstitutions, databaseInstitutions)
	}

	fmt.Printf("Imported %d", len(importedInstitutions))

	// transpose the remaining institutions into database institutions
	transposedInstitutions := make([]DatabaseInstitution, 0)
	for _, webInstitution := range webInstitutions {
		var row DatabaseInstitution
		row.Name = webInstitution.Name
		row.InternationalName = webInstitution.Name
		row.Alias = ""
		row.Domain = webInstitution.Domains[0]
		row.ImageUrl = webInstitution.ImageUrl
		row.CountryIsoCode = webInstitution.AlphaTwoCode
		row.IsSupported = false

		transposedInstitutions = append(transposedInstitutions, row)
	}

	// write the transposed institutions to a json file
	file, err := os.Create("unknownInstitutions.json")
	catch(err, 1, "Failed to create file")
	defer file.Close()

	json.NewEncoder(file).Encode(transposedInstitutions)

	// write the imported institutions to a json file
	file, err = os.Create("knownInstitutions.json")
	catch(err, 1, "Failed to create file")

	defer file.Close()

	json.NewEncoder(file).Encode(importedInstitutions)
}

func findAndProcess(index string, webInstitutions map[string]WebInstitution, databaseInstitutions map[string]DatabaseInstitution) {
	var found = false
	for _, databaseInstitution := range databaseInstitutions {
		for _, domain := range webInstitutions[index].Domains {
			if domain == databaseInstitution.Domain {
				found = true
				importedInstitutions[databaseInstitution.InternationalName] = &databaseInstitution
				delete(webInstitutions, webInstitutions[index].Name)
				break
			}
		}
	}

	if !found {
		// find the root domain, meaning the domain with the least amount of dots
		var rootDomain = ""
		var rootDomainLength = 1000
		for _, domain := range webInstitutions[index].Domains {
			if len(domain) < rootDomainLength {
				rootDomain = domain
				rootDomainLength = len(domain)
			}
		}

		// fetch logo
		tmp := webInstitutions[index]
		tmp.ImageUrl = fetchLogo(rootDomain)
		webInstitutions[index] = tmp
	}
}
