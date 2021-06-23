package bar

import "testing"

func TestInheritance(t *testing.T) {
	b := NewBar()
	err := b.Fooo()
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

}