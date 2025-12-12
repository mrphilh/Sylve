package iscsi

import (
	"strings"

	uclConfigsInterfaces "github.com/alchemillahq/sylve/internal/interfaces/configs/ucl"
)

//"github.com/alchemillahq/sylve/pkg/system"

func buildConfig(contexts map[string][]uclConfigsInterfaces.UclConfigInterface) string {
	config := strings.Builder{}

	//Order of config
	// global settings
	if gs, ok := contexts["global"]; ok {
		config.WriteString(gs[0].AsUcl(0))
	}
	// auth-group
	if ag, ok := contexts["auth-group"]; ok {
		config.WriteString("auth-group {\n")
		for i := 0; i < len(ag); i++ {
			config.WriteString(ag[i].AsUcl(1))
		}
		config.WriteString("}\n")
	}
	// portal-group
	if pg, ok := contexts["portal-group"]; ok {
		config.WriteString("portal-group {\n")
		for i := 0; i < len(pg); i++ {
			config.WriteString(pg[i].AsUcl(1))
		}
		config.WriteString("}\n")
	}
	// lun
	if l, ok := contexts["lun"]; ok {
		config.WriteString("lun {\n")
		for i := 0; i < len(l); i++ {
			config.WriteString(l[i].AsUcl(1))
		}
		config.WriteString("}\n")
	}
	// target
	if t, ok := contexts["target"]; ok {
		config.WriteString("target {\n")
		for i := 0; i < len(t); i++ {
			config.WriteString(t[i].AsUcl(1))
		}
		config.WriteString("}\n")
	}
	// TODO: controller
	//if c, ok := contexts["controller"]; ok {
	//	config.WriteString("target {\n")
	//	config.WriteString(c.AsUcl(2))
	//	config.WriteString("}\n")
	//}

	return config.String()
}
func (s *Service) WriteConfig(reload bool) error {
	// Load config data from db
	contexts := make(map[string][]uclConfigsInterfaces.UclConfigInterface)

	// global settings
	globalSettings, err := s.GetGlobalSettings()
	if err != nil {
		return err
	}

	contexts["global"] = append(contexts["global"], &globalSettings[0])

	// auth-group
	authGroups, err := s.GetAuthGroups()
	if err != nil {
		return err
	}
	for i := 0; i < len(authGroups); i++ {
		contexts["auth-group"] = append(contexts["auth-group"], &authGroups[i])
	}

	// portal-group
	portalGroups, err := s.GetPortalGroups()
	if err != nil {
		return err
	}
	for i := 0; i < len(portalGroups); i++ {
		contexts["portal-group"] = append(contexts["portal-group"], &portalGroups[i])
	}

	// lun
	luns, err := s.GetLuns()
	if err != nil {
		return err
	}
	for i := 0; i < len(luns); i++ {
		contexts["lun"] = append(contexts["lun"], &luns[i])
	}

	// target
	targets, err := s.GetTargets()
	if err != nil {
		return err
	}
	for i := 0; i < len(luns); i++ {
		contexts["target"] = append(contexts["target"], &targets[i])
	}
	// TODO: controller

	//create config string
	//_configFile = buildConfig(contexts)
	//write config to disk

	//reload ctld
	/*
		if err := system.ServiceAction("ctld", "reload"); err != nil {
			return fmt.Errorf("filaed to reload ctld service: %w", err)
		}
	*/
	return nil

}
