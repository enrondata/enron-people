package enronpeople

import (
	"regexp"
	"testing"
)

var peopleTests = []struct {
	userName     string
	emailAddress string
}{
	{"skilling-j", "jeffery.skilling@enron.com"},
}

var rxSpace = regexp.MustCompile(`[\s\t]`)

func TestPeople(t *testing.T) {
	peopleSet := NewPeopleSet()
	t.Logf("user set count [%d/%d]", peopleSet.Count(), len(UserNamesCalo()))
	for _, tt := range peopleTests {
		try := peopleSet.GetByEmail(tt.emailAddress)
		if try == nil {
			t.Errorf("PeopleSet.GetByEmail(\"%s\") Error: user not found", tt.emailAddress)
		} else if tt.userName != try.UserName {
			t.Errorf("PeopleSet.GetByEmail(\"%s\") Error: want [%s], got [%s]",
				tt.emailAddress, tt.userName, try.UserName)
		}
	}
}
