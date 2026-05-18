package enronpeople

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/grokify/goauth/scim"
)

//go:embed data/edo_enron-custodians-data.json
var custodiansFS embed.FS

// SchemaPerson represents a person in Schema.org format (from EDO data).
type SchemaPerson struct {
	Context       string           `json:"@context"`
	Type          string           `json:"@type"`
	AlternateName string           `json:"alternateName"` // maildir folder name
	GivenName     string           `json:"givenName"`
	FamilyName    string           `json:"familyName"`
	Email         string           `json:"email"`
	Name          string           `json:"name"`
	JobTitle      string           `json:"jobTitle,omitempty"`
	Affiliation   SchemaAffilation `json:"affiliation,omitempty"`
	Maildirs      []string         `json:"maildirs,omitempty"`
}

// SchemaAffilation represents affiliation in Schema.org format.
type SchemaAffilation struct {
	LegalName string `json:"legalName"`
}

// LoadCustodiansJSON loads custodians from the embedded JSON file.
func LoadCustodiansJSON() ([]SchemaPerson, error) {
	data, err := custodiansFS.ReadFile("data/edo_enron-custodians-data.json")
	if err != nil {
		return nil, fmt.Errorf("read custodians file: %w", err)
	}

	var people []SchemaPerson
	if err := json.Unmarshal(data, &people); err != nil {
		return nil, fmt.Errorf("unmarshal custodians: %w", err)
	}

	return people, nil
}

// NewCustodiansUserSet creates a SCIM UserSet from the custodians JSON data.
func NewCustodiansUserSet() (scim.UserSet, error) {
	people, err := LoadCustodiansJSON()
	if err != nil {
		return scim.UserSet{}, err
	}

	users := make([]scim.User, 0, len(people))
	for _, p := range people {
		user := SchemaPersonToSCIMUser(p)
		users = append(users, user)
	}

	return scim.UserSet{Users: users}, nil
}

// SchemaPersonToSCIMUser converts a Schema.org Person to a SCIM User.
func SchemaPersonToSCIMUser(p SchemaPerson) scim.User {
	emails := []scim.Item{}

	// Add primary email
	if p.Email != "" {
		emails = append(emails, scim.Item{
			Value:   strings.ToLower(p.Email),
			Type:    TypeWork,
			Primary: true,
		})
	}

	// Generate common email variations
	if p.GivenName != "" && p.FamilyName != "" {
		given := strings.ToLower(p.GivenName)
		family := strings.ToLower(p.FamilyName)

		variations := []string{
			fmt.Sprintf("%s.%s@enron.com", given, family),
			fmt.Sprintf("%s_%s@enron.com", given, family),
			fmt.Sprintf("%s%s@enron.com", given[:1], family),
			fmt.Sprintf("%s@enron.com", family),
		}

		// Add maildir-based variations
		if p.AlternateName != "" {
			parts := strings.Split(p.AlternateName, "-")
			if len(parts) == 2 {
				// e.g., "skilling-j" -> jskilling@enron.com
				variations = append(variations,
					fmt.Sprintf("%s%s@enron.com", parts[1], parts[0]),
					fmt.Sprintf("%s.%s@enron.com", parts[1], parts[0]),
				)
			}
		}

		// Add unique variations
		seen := make(map[string]bool)
		if p.Email != "" {
			seen[strings.ToLower(p.Email)] = true
		}
		for _, v := range variations {
			if !seen[v] {
				seen[v] = true
				emails = append(emails, scim.Item{
					Value:   v,
					Type:    TypeWork,
					Primary: false,
				})
			}
		}
	}

	department := ""
	if p.Affiliation.LegalName != "" {
		department = p.Affiliation.LegalName
	}

	return scim.User{
		Active:      true,
		DisplayName: p.Name,
		Emails:      emails,
		Name: scim.Name{
			GivenName:  p.GivenName,
			FamilyName: p.FamilyName,
		},
		Title:    p.JobTitle,
		UserName: p.AlternateName,
		UserType: scim.UserTypeEmployee,
		Groups:   []scim.Group{{Display: department}},
	}
}
