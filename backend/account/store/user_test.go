package store

import "testing"

func sum(a, b int) int {
	return a + b
}

func Test_CreateUser(t *testing.T) {
	a := []struct {
		expect  int
		a       int
		b       int
		message string
	}{
		{20, 1, 2, "should fail"},
		{19, 3, 2, "should fail"},
	}

	for _, v := range a {
		if sum(v.a, v.b) != v.expect {
			t.Errorf("expect %d, got %d", v.expect, v.a)
		}
	}
}
