package iscsiModels

import (
	"testing"
)

func TestTargetAsUcl(t *testing.T) {
	alias := "target alias"
	var tests = []struct {
		name     string
		tg       *Target
		expected string
	}{
		{
			"Model Defaults",
			&Target{Name: "iqn.2012-06.com.example:target0", AuthGroup: AuthGroup{Name: "ag0"}, PortalGroup: PortalGroup{Name: "ag0"}, Luns: []Lun{Lun{Name: "lun0", Path: "/test/path", Size: "1GB"}}},
			"\"iqn.2012-06.com.example:target0\" {\n" +
				"\tauth-group = ag0\n" +
				"\tportal-group = pg0\n" +
				"\tlun = {\n" +
				"\t\t0 = lun0\n" +
				"\t}\n" +
				"}\n",
		},
		{
			"Model with alias",
			&Target{
				Name:        "iqn.2012-06.com.example:target0",
				Alias:       &alias,
				AuthGroup:   AuthGroup{Name: "ag0"},
				PortalGroup: PortalGroup{Name: "ag0"},
				Luns: []Lun{
					Lun{Name: "lun0", Path: "/test/path", Size: "1GB"},
				},
			},
			"\"iqn.2012-06.com.example:target0\" {\n" +
				"\talias = \"target alias\"\n" +
				"\tauth-group = ag0\n" +
				"\tportal-group = pg0\n" +
				"\tlun = {\n" +
				"\t\t0 = lun0\n" +
				"\t}\n" +
				"}\n",
		},
		{
			"Model with 3 luns",
			&Target{
				Name:        "iqn.2012-06.com.example:target0",
				Alias:       &alias,
				AuthGroup:   AuthGroup{Name: "ag0"},
				PortalGroup: PortalGroup{Name: "ag0"},
				Luns: []Lun{
					Lun{Name: "lun0", Path: "/test/path0", Size: "1GB"},
					Lun{Name: "lun1", Path: "/test/path1", Size: "2GB"},
					Lun{Name: "lun2", Path: "/test/path2", Size: "3GB"},
				},
			},
			"\"iqn.2012-06.com.example:target0\" {\n" +
				"\talias = \"target alias\"\n" +
				"\tauth-group = ag0\n" +
				"\tportal-group = pg0\n" +
				"\tlun = {\n" +
				"\t\t0 = lun0\n" +
				"\t\t1 = lun1\n" +
				"\t\t2 = lun2\n" +
				"\t}\n" +
				"}\n",
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rendered := tt.tg.AsUcl(0)
			if rendered != tt.expected {
				t.Errorf("\n------got------\n%s\n------expected------\n%s\n------\n", rendered, tt.expected)
			}
		})
	}
}
