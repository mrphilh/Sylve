package iscsi

import (
	"testing"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
	uclConfigsInterfaces "github.com/alchemillahq/sylve/internal/interfaces/configs/ucl"
)

func Test_buildConfig(t *testing.T) {
	ag0 := iscsiModels.AuthGroup{Name: "ag0", AuthType: "none"}
	ag1 := iscsiModels.AuthGroup{Name: "ag1", AuthType: "none"}
	pg0 := iscsiModels.PortalGroup{Name: "pg0", DiscoveryAuthGroup: "default"}
	pg1 := iscsiModels.PortalGroup{Name: "pg1", DiscoveryAuthGroup: "default"}

	lun0 := iscsiModels.Lun{
		Name: "lun0",
		Path: "/test/path0",
		Size: "1GB",
	}
	lun1 := iscsiModels.Lun{
		Name: "lun1",
		Path: "/test/path1",
		Size: "1GB",
	}
	lun2 := iscsiModels.Lun{
		Name: "lun2",
		Path: "/test/path2",
		Size: "1GB",
	}
	target0 := iscsiModels.Target{
		Name:        "iqn.2012-06.com.example:target0",
		AuthGroup:   ag0,
		PortalGroup: pg0,
		Luns: []iscsiModels.Lun{
			lun0,
		},
	}
	target1 := iscsiModels.Target{
		Name:        "iqn.2012-06.com.example:target1",
		AuthGroup:   ag1,
		PortalGroup: pg1,
		Luns: []iscsiModels.Lun{
			lun1,
			lun2,
		},
	}

	context0 := make(map[string][]uclConfigsInterfaces.UclConfigInterface)
	context0["global"] = append(context0["global"], &iscsiModels.GlobalSetting{})
	context0["auth-group"] = append(context0["auth-group"], &ag0)
	context0["portal-group"] = append(context0["portal-group"], &pg0)
	context0["lun"] = append(context0["lun"], &lun0)
	context0["target"] = append(context0["target"], &target0)

	context1 := make(map[string][]uclConfigsInterfaces.UclConfigInterface)
	context1["global"] = append(context1["global"], &iscsiModels.GlobalSetting{})
	context1["auth-group"] = append(context1["auth-group"], &ag0)
	context1["auth-group"] = append(context1["auth-group"], &ag1)
	context1["portal-group"] = append(context1["portal-group"], &pg0)
	context1["portal-group"] = append(context1["portal-group"], &pg1)
	context1["lun"] = append(context1["lun"], &lun0)
	context1["lun"] = append(context1["lun"], &lun1)
	context1["lun"] = append(context1["lun"], &lun2)
	context1["target"] = append(context1["target"], &target0)
	context1["target"] = append(context1["target"], &target1)

	var tests = []struct {
		name     string
		contexts map[string][]uclConfigsInterfaces.UclConfigInterface
		expected string
	}{
		{
			"Empty Contexts",
			context0,
			"debug = 0\n" +
				"auth-group {\n" +
				"\tag0 {\n" +
				"\t\tauth-type = none\n" +
				"\t}\n" +
				"}\n" +
				"portal-group {\n" +
				"\tpg0 {\n" +
				"\t\tdiscovery-auth-group = default\n" +
				"\t}\n" +
				"}\n" +
				"lun {\n" +
				"\tlun0 {\n" +
				"\t\tpath = /test/path0\n" +
				"\t\tsize = 1GB\n" +
				"\t}\n" +
				"}\n" +
				"target {\n" +
				"\t\"iqn.2012-06.com.example:target0\" {\n" +
				"\t\tauth-group = ag0\n" +
				"\t\tportal-group = pg0\n" +
				"\t\tlun = {\n" +
				"\t\t\t0 = lun0\n" +
				"\t\t}\n" +
				"\t}\n" +
				"}\n",
		},
		{
			"Multiple elements per context",
			context1,
			"debug = 0\n" +
				"auth-group {\n" +
				"\tag0 {\n" +
				"\t\tauth-type = none\n" +
				"\t}\n" +
				"\tag1 {\n" +
				"\t\tauth-type = none\n" +
				"\t}\n" +
				"}\n" +
				"portal-group {\n" +
				"\tpg0 {\n" +
				"\t\tdiscovery-auth-group = default\n" +
				"\t}\n" +
				"\tpg1 {\n" +
				"\t\tdiscovery-auth-group = default\n" +
				"\t}\n" +
				"}\n" +
				"lun {\n" +
				"\tlun0 {\n" +
				"\t\tpath = /test/path0\n" +
				"\t\tsize = 1GB\n" +
				"\t}\n" +
				"\tlun1 {\n" +
				"\t\tpath = /test/path1\n" +
				"\t\tsize = 1GB\n" +
				"\t}\n" +
				"\tlun2 {\n" +
				"\t\tpath = /test/path2\n" +
				"\t\tsize = 1GB\n" +
				"\t}\n" +
				"}\n" +
				"target {\n" +
				"\t\"iqn.2012-06.com.example:target0\" {\n" +
				"\t\tauth-group = ag0\n" +
				"\t\tportal-group = pg0\n" +
				"\t\tlun = {\n" +
				"\t\t\t0 = lun0\n" +
				"\t\t}\n" +
				"\t}\n" +
				"\t\"iqn.2012-06.com.example:target1\" {\n" +
				"\t\tauth-group = ag1\n" +
				"\t\tportal-group = pg1\n" +
				"\t\tlun = {\n" +
				"\t\t\t0 = lun1\n" +
				"\t\t\t1 = lun2\n" +
				"\t\t}\n" +
				"\t}\n" +
				"}\n",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			rendered := buildConfig(tt.contexts)
			if rendered != tt.expected {
				t.Errorf("\n------got------\n%s\n------expected------\n%s\n------\n", rendered, tt.expected)
			}
		})
	}

}
