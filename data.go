package enronpeople

import (
	"github.com/grokify/oauth2more/scim"
	"github.com/grokify/simplego/time/timezone"
)

const TypeWork = "work"

// Exchange Legacy DN: /O=ENRON/OU=NA/CN=RECIPIENTS/CN=Jskillin
// Full x500 address: x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Jskillin

func NewPeopleSet() scim.UserSet {
	return scim.UserSet{
		Users: people}
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
		DisplayName: "Harry Arora",
		Emails: []scim.Item{
			{
				Value: "harry.arora@enron.com",
				Type:  TypeWork, Primary: true},
			{
				Value: "harpreet.arora@enron.com",
				Type:  TypeWork, Primary: false}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Harora",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			GivenName:  "Harpreet",
			MiddleName: "S.",
			FamilyName: "Arora"},
		NickName: "Harry",
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "VP, Trading",
		UserName: "arora-h",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Jeffrey Keith Skilling",
		Emails: []scim.Item{{
			Value: "jeffery.skilling@enron.com",
			Type:  TypeWork, Primary: true}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Jskillin",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			GivenName:  "Jeffrey",
			MiddleName: "Keith",
			FamilyName: "Skilling"},
		NickName:   "Jeff",
		ProfileURL: "https://en.wikipedia.org/wiki/Jeffrey_Skilling",
		Timezone:   timezone.TimezoneAmericaChicago,
		Title:      "CEO",
		UserName:   "skilling-j",
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
		ExternalID:  "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Klay",
		Emails: []scim.Item{{
			Value: "kenneth.lay@enron.com",
			Type:  TypeWork, Primary: true}},
		Groups: []scim.Group{{Display: "Enron Corp."}},
		Locale: "en-US",
		Name: scim.Name{
			GivenName:  "Kenneth",
			MiddleName: "Lee",
			FamilyName: "Lay"},
		ProfileURL: "https://en.wikipedia.org/wiki/Kenneth_Lay",
		Timezone:   timezone.TimezoneAmericaChicago,
		Title:      "CEO, Chairman, Founder",
		UserName:   "lay-k",
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
