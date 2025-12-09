package iscsiModels

import (
	"testing"
)

/*
*       portal-group {
	       pg0 {
		       discovery-auth-group = no-authentication
		       listen =	[
			       0.0.0.0:3260,
			       "[::]:3260",
			       "[fe80::be:ef]:3261"
		       ]
	       }
       }
*/

func TestPortalGroupAsUcl(t *testing.T) {
	addr := "0.0.0.0"
	var tests = []struct {
		name     string
		pg       *PortalGroup
		expected string
	}{
		{
			"Model Defaults",
			&PortalGroup{Name: "pg0", DiscoveryAuthGroup: "default"},
			"pg0 {\n" +
				"\tdiscovery-auth-group = default\n" +
				"}\n",
		},
		{
			"debug parameter override",
			&PortalGroup{Name: "pg0", DiscoveryAuthGroup: "no-authentication", Listen: &addr},
			"pg0 {\n" +
				"\tdiscovery-auth-group = no-authentication\n" +
				"\tlisten = 0.0.0.0\n" +
				"}\n",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rendered := tt.pg.AsUcl(0)
			if rendered != tt.expected {
				t.Errorf("\n------got------\n%s\n------expected------\n%s\n------\n", rendered, tt.expected)
			}
		})
	}
}
