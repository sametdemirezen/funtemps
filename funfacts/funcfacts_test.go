package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string 
		want []string  
	}

	tests := []test{
		{input: "terra", want: []string{
		"Høyeste temperatur målt på Jordens overflate.", 
		"Laveste temperatur målt på Jordens overflate.", 
		"Temperatur i Jordens indre kjerne"}},

		{input: "sun", want: []string{
		"Temperatur i Solens kjerne.", 
		"Temperatur på ytre lag av Solen."}},

		{input: "luna", want: []string{
		"Temperatur på Månens overflate om natten.", 
		"Temperatur på Månens overflate om dagen."}},
		
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
