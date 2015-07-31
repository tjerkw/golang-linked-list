package tjerkw

import "testing"

func TestSet(t *testing.T) {

	s := NewSet()
	if s.Contains(3) {
		t.Fatal("Should not be found")
	}
	s.Add(4)
	if s.Contains(3) {
		t.Fatal("Should not be found")
	}
	if !s.Contains(4) {
		t.Fatal("Should be found")
	}
	s.Add(2)
	s.Add(52)
	err := s.Add(4)
	if err == nil {
		t.Fatal("Err should not be nil")
	}
}
