package random

import(
	"testing"
)

func TestRandomString(t *testing.T){
	expected := "a"
	got := RandomString(1)

	if len(got) != len(expected){
		t.Errorf("Expected %v, Got %v", expected, got)
	}
}

func TestRandomNumber(t *testing.T){
	expected := []int{0, 1}
	got := RandomNumber(0, 1)

	if got != expected[0] && got != expected[1]{
		t.Errorf("Expected %v, Got, %v", expected, got)
	}
}