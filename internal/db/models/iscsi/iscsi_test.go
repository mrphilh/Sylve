package iscsiModels

import (
	"testing"
)

func TestGlobalSettingsAsUcl(t *testing.T) {
	var tests = []struct {
		name     string
		gs       *GlobalSetting
		expected string
	}{
		{
			"Model Defaults",
			&GlobalSetting{},
			"debug = 0\n",
		},
		{
			"debug parameter override",
			&GlobalSetting{Debug: 2},
			"debug = 2\n",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rendered := tt.gs.AsUcl(0)
			if rendered != tt.expected {
				t.Errorf("got\n------\n%s\n------\nexpected:\n------\n%s\n------\n", rendered, tt.expected)
			}
		})
	}
}
