package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want []string  // her må du skrive riktig type for returverdien
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{input: "terra", want: []string{"Høyeste temperatur målt på Jordens overflate.", "Laveste temperatur målt på Jordens overflate.", "Temperatur i Jordens indre kjerne"}},
		{input: "sun", want: []string{"Temperatur i Solens kjerne.", "Temperatur på ytre lag av Solen."}},
		{input: "luna", want: []string{"Temperatur på Månens overflate om natten.", "Temperatur på Månens overflate om dagen."}},
		{input: "invalid", want: []string{}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
