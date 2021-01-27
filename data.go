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
		Emails: []scim.Item{
			{
				Value: "david.delainey@enron.com",
				Type:  TypeWork, Primary: true},
			{
				Value: "w..delainey@enron.com",
				Type:  TypeWork, Primary: false}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Ddelain2",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			GivenName:  "David",
			MiddleName: "W.",
			FamilyName: "Delainey"},
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "CEO, Enron North America and Enron Energy Services",
		UserName: "delainey-d",
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
		Emails: []scim.Item{{
			Value: "james.derrick@enron.com",
			Type:  TypeWork, Primary: true}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=JDERRIC",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			GivenName:       "James",
			FamilyName:      "Derrick",
			HonorificSuffix: "Jr."},
		NickName: "Jim",
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "Executive Vice President and General Counsel",
		UserName: "derrick-j",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "John Arnold",
		Emails: []scim.Item{{
			Value: "john.arnold@enron.com",
			Type:  TypeWork, Primary: true}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Jarnold",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			GivenName:  "John",
			FamilyName: "Arnold"},
		Timezone: timezone.TimezoneAmericaChicago,
		Title:    "Vice President",
		UserName: "arnold-j",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Kenneth Lee Lay",
		Emails: []scim.Item{{
			Value: "kenneth.lay@enron.com",
			Type:  TypeWork, Primary: true}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Klay",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			Formatted:  "Kenneth Lee Lay",
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
		DisplayName: "Phillip Allen",
		Emails: []scim.Item{{
			Value: "phillip.allen@enron.com",
			Type:  TypeWork, Primary: true}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Pallen",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			Formatted:  "Phillip K. Allen",
			GivenName:  "Phillip",
			MiddleName: "K.",
			FamilyName: "Allen"},
		Timezone: timezone.TimezoneAmericaChicago,
		UserName: "allen-p",
		UserType: scim.UserTypeEmployee},
	{
		Active:      true,
		DisplayName: "Robert Badeer",
		Emails: []scim.Item{{
			Value: "robert.badeer@enron.com",
			Type:  TypeWork, Primary: true}},
		ExternalID: "x500:/O=ENRON/OU=NA/CN=RECIPIENTS/CN=Rbadeer",
		Groups:     []scim.Group{{Display: "Enron Corp."}},
		Locale:     "en-US",
		Name: scim.Name{
			GivenName:  "Robert",
			FamilyName: "Badeer"},
		NickName: "Bob",
		Timezone: timezone.TimezoneAmericaLosAngeles,
		Title:    "Director, Trading, West Coast Trading Desk, Portland, OR",
		UserName: "badeer-r",
		UserType: scim.UserTypeEmployee},
}
