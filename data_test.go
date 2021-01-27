package enronpeople

import (
	"testing"
)

var peopleTests = []struct {
	externalID   string
	emailAddress string
}{
	{"skilling-j", "jeffery.skilling@enron.com"},
}

func TestPeople(t *testing.T) {
	peopleSet := NewPeopleSet()
	t.Logf("user set count [%d]", peopleSet.Count())
	for _, tt := range peopleTests {
		try := peopleSet.GetByEmail(tt.emailAddress)
		if try == nil {
			t.Errorf("PeopleSet.GetByEmail(\"%s\") Error: user not found",
				tt.emailAddress)
		}
		if tt.externalID != try.ExternalID {
			t.Errorf("PeopleSet.GetByEmail(\"%s\") Error: want [%s], got [%s]",
				tt.emailAddress, tt.externalID, try.ExternalID)
		}
	}
}
