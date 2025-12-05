package iscsiModels

import (
	"testing"
)

func TestLunAsUcl(t *testing.T) {
	var tests = []struct {
		name     string
		l        *Lun
		expected string
	}{
		{
			"Model Defaults",
			&Lun{Name: "lun0", Path: "/test/path", Size: "1GB"},
			"lun0 {\n" +
				"\tpath = /test/path\n" +
				"\tsize = 1GB\n" +
				"}\n",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rendered := tt.l.AsUcl(0)
			if rendered != tt.expected {
				t.Errorf("got\n------\n%s\n------\nexpected:\n------\n%s\n------\n", rendered, tt.expected)
			}
		})
	}
}
