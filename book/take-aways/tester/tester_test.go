package tester

import (
	"testing"
)

func TestTeam(t *testing.T) {
	tm := NewTeams()
	tm.AddPlayer(Person{Name: "A", Age: 10, Address: Address{Street: "123", City: "New York", Zip: "12345"}})
	tm.AddPlayer(Employee{Name: "B", Age: 10, Address: Address{Street: "123", City: "New York", Zip: "12345"}})
	tm.AddPlayer(Person{Name: "C", Age: 10, Address: Address{Street: "123", City: "New York", Zip: "12345"}})

	if len(tm.Players) != 3 {
		t.Errorf("Expected 3 persons, got %d", len(tm.Players))
	}
}
