package main

import "testing"

func TestIKnowMoreTestsAreNeeded(t *testing.T) {
	expected := "More tests please!"
	if expected != "More tests please!" {
		t.Errorf("%v, I know!", expected)
	}
}
