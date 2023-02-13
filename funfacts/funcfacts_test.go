package funfacts

import (
	"reflect"
	"testing"
	   )

func TestGetFunFacts(t *testing.T) {
	tests := []struct {
		input string // her må du skrive riktig type for input
		want  []string // her må du skrive riktig type for returverdien
	}{
	{input: "sun" , want: facts.Sun},
	{input: "luna" , want: facts.Luna},
	{input: "terra" , want: facts.Terra},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}

}