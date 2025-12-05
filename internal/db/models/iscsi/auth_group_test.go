package iscsiModels

import (
	"testing"
)

func TestAuthGroupAsUcl(t *testing.T) {
	var tests = []struct {
		name     string
		ag       *AuthGroup
		expected string
	}{
		{
			"Model Defaults",
			&AuthGroup{Name: "ag0", AuthType: "none"},
			"ag0 {\n" +
				"\tauth-type = none\n" +
				"}\n",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rendered := tt.ag.AsUcl(0)
			if rendered != tt.expected {
				t.Errorf("got\n------\n%s\n------\nexpected:\n------\n%s\n------\n", rendered, tt.expected)
			}
		})
	}
}
