package enronpeople

import (
	"fmt"
	"strings"

	"github.com/grokify/oauth2more/scim"
	"github.com/grokify/simplego/time/timezone"
)

func GetPeople() []scim.User {
	return people
}

type PeopleSet struct {
	people []scim.User
}

func NewPeopleSet() PeopleSet {
	return PeopleSet{
		people: people}
}

func (set *PeopleSet) Count() int {
	return len(set.people)
}

func (set *PeopleSet) GetByEmail(emailAddress string) *scim.User {
	emailAddress = strings.TrimSpace(emailAddress)
	matches := []scim.User{}
	for _, person := range set.people {
		if emailAddress == person.UserName {
			matches = append(matches, person)
		}
	}
	if len(matches) > 1 {
		panic(fmt.Sprintf("non-unique email address in set [%s]", emailAddress))
	} else if len(matches) == 1 {
		return &matches[0]
	}
	return nil
}

var people = []scim.User{
	{
		Active:      true,
		DisplayName: "David Delainey",
		ExternalID:  "delainey-d",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "David",
			FamilyName: "Delainey"},
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "CEO, Enron North America and Enron Energy Services",
		UserName: "david.delainey@enron.com",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Harpreet S. Arora",
		ExternalID:  "arora-h",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "Harpreet",
			MiddleName: "S.",
			FamilyName: "Arora"},
		NickName: "Harry",
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "VP, Trading",
		UserName: "harpreet.arora@enron.com",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Jeffrey Keith Skilling",
		ExternalID:  "skilling-j",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "Jeffrey",
			MiddleName: "Keith",
			FamilyName: "Skilling"},
		NickName:   "Jeff",
		ProfileURL: "https://en.wikipedia.org/wiki/Jeffrey_Skilling",
		Timezone:   timezone.TimezoneAmericaChicago,
		Title:      "CEO",
		UserName:   "jeffery.skilling@enron.com",
		UserType:   scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Jim Derrick",
		ExternalID:  "derrick-j",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "James",
			FamilyName: "Derrick"},
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "Executive Vice President and General Counsel",
		UserName: "james.derrick@enron.com",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Kenneth Lee Lay",
		ExternalID:  "derrick-j",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "Kenneth",
			MiddleName: "Lee",
			FamilyName: "Lay"},
		ProfileURL: "https://en.wikipedia.org/wiki/Kenneth_Lay",
		Timezone:   timezone.TimezoneAmericaChicago,
		Title:      "CEO, Chairman, Founder",
		UserName:   "kenneth.lay@enron.com",
		UserType:   scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Philip Allen",
		ExternalID:  "allen-p",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "Philip",
			FamilyName: "Allen"},
		Timezone: timezone.TimezoneAmericaChicago,
		UserName: "philip.allen@enron.com",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Robert Badeer",
		ExternalID:  "badeer-r",
		Groups:      []scim.Group{{Display: "Enron Corp."}},
		Locale:      "en-US",
		Name: scim.Name{
			GivenName:  "Robert",
			FamilyName: "Badeer"},
		NickName: "Bob",
		Timezone: timezone.TimezoneAmericaLosAngeles,
		Title:    "Director, Trading, West Coast Trading Desk, Portland, OR",
		UserName: "robert.badeer@enron.com",
		UserType: scim.UserTypeEmployee},
}
