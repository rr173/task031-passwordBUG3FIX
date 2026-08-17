package password

import "testing"

func TestBug3MaxLengthIsInclusive(t *testing.T) {
	r := Evaluate("abc", Policy{MaxLength: 3})
	if !r.OK {
		t.Fatalf("password at maxLength must pass, got violations=%v", r.Violations)
	}
}
