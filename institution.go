package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/huandu/go-sqlbuilder"
)

type DatabaseInstitution struct {
	ID                string `db:"id" json:"id"`
	Name              string `db:"name" json:"name"`
	InternationalName string `db:"internationalName" json:"internationalName"`
	Alias             string `db:"alias" json:"alias"`
	Domain            string `db:"domain" json:"domain"`
	ImageUrl          string `db:"imageUrl" json:"imageUrl"`
	CountryIsoCode    string `db:"countryIsoCode" json:"countryIsoCode"`

	// optional fields
	IsSupported bool `db:"isSupported" json:"isSupported"`
}

var institutionStruct = sqlbuilder.NewStruct(new(DatabaseInstitution))

func GetInstitutionsFromDatabase() map[string]DatabaseInstitution {
	dbURI, isSet := os.LookupEnv("DATABASE")
	if !isSet {
		panic(errors.New("Missing database env variable"))
	}

	_sql, err := sql.Open("postgres", dbURI)
	catch(err, 1, "Failed to connect to database")

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("id", "name", "\"internationalName\"", "alias", "domain", "\"imageUrl\"", "\"countryIsoCode\"", "\"isSupported\"").
		From("students.institution")

	query, args := sb.Build()
	rows, err := _sql.Query(query, args...)
	catch(err, 1, "An error occurred while fetching institution data")

	institutions := make(map[string]DatabaseInstitution)
	for rows.Next() {

		var row DatabaseInstitution

		err = rows.Scan(institutionStruct.Addr(&row)...)
		catch(err, 1, "Could not copy data to memory addr for "+row.Name)

		institutions[row.InternationalName] = row
	}

	return institutions
}

type WebInstitution struct {
	Name         string   `json:"name"`
	Domains      []string `json:"domains"`
	WebPages     []string `json:"web_pages"`
	Country      string   `json:"country"`
	AlphaTwoCode string   `json:"alpha_two_code"`
	ImageUrl     string
}

func GetInstitutionsFromWeb() map[string]WebInstitution {
	res, err := http.Get("https://github.com/UniStudents/university-domains-list/raw/master/world_universities_and_domains.json")
	catch(err, 1, "Could not fetch institutions from Github repository")

	unparsedInstitutions := make([]WebInstitution, 0)
	err = json.NewDecoder(res.Body).Decode(&unparsedInstitutions)
	catch(err, 1, "Could not parse JSON data from Github repository")

	institutions := make(map[string]WebInstitution)

	for _, institution := range unparsedInstitutions {
		institutions[institution.Name] = institution
	}

	return institutions

}

type Country struct {
}
