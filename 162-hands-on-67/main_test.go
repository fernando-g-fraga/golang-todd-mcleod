package main

import "testing"

func TestGetUser(t *testing.T) {

	md := &MockDatastore{
		Users: map[int]User{
			2: {ID: 2, first: "Marcos"},
		},
	}

	s := &Service{ds: md}

	u, err := s.GetUser(2)

	if err != nil {
		t.Errorf("got error: %v", err)
	}

	if u.first != "Marcos" {
		t.Errorf("Got: %s Want: %s", u.first, "Marcos")
	}
}
