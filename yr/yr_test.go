package yr

import (
	"reflect"
	"testing"
)

func TestConvTextSlutt(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", want: "er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Viktor Kalhovd"},
	}

	for _, tc := range tests {
		got := convTextSlutt(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("Expected: %v, got: %v", tc.want, got)
		}
	}
}
