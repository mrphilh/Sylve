package iscsi

import (
	"fmt"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
)

func (s *Service) GetPortalGroups() ([]iscsiModels.PortalGroup, error) {
	var portalGroups []iscsiModels.PortalGroup
	if err := s.DB.Find(&portalGroups).Error; err != nil {
		return nil, fmt.Errorf("failed_to_get_portal_groups: %w", err)
	}
	return portalGroups, nil
}

func (s *Service) CreatePortalGroup(name string, discoveryAuthGroup string, listen string) error {
	if err := s.DB.Where("name = ?", name).First(&iscsiModels.PortalGroup{}).Error; err == nil {
		return fmt.Errorf("portal_group_with_name_exists")
	}

	portalGroup := iscsiModels.PortalGroup{
		Name:               name,
		DiscoveryAuthGroup: discoveryAuthGroup,
		Listen:             &listen,
	}

	if err := s.DB.Create(&portalGroup).Error; err != nil {
		return fmt.Errorf("failed_to_create_portal_group: %w", err)
	}

	return s.WriteConfig(true)
}

func (s *Service) UpdatePortalGroup(id uint, name string, discoveryAuthGroup string, listen string) error {
	var portalGroup iscsiModels.PortalGroup
	if err := s.DB.First(&portalGroup, id).Error; err != nil {
		return fmt.Errorf("portal_group_not_found %w", err)
	}

	if name != portalGroup.Name {
		var count int64
		if err := s.DB.Model(&iscsiModels.PortalGroup{}).Where("name = ? AND id != ?", name, id).Count(&count).Error; err != nil {
			return fmt.Errorf("failed_to_check_name_confict: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("portal_group_with_name_exists")
		}

		tx := s.DB.Begin()

		portalGroup.Name = name
		portalGroup.DiscoveryAuthGroup = discoveryAuthGroup
		portalGroup.Listen = &listen

		if err := tx.Save(&portalGroup).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed_to_update_portal_group_fields: %w", err)
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed_to_commit_transaction: %w", err)
		}
	}
	return s.WriteConfig(true)
}

func (s *Service) DeletePortalGroup(id uint) error {
	var portalGroup iscsiModels.PortalGroup
	if err := s.DB.First(&portalGroup, id).Error; err != nil {
		return fmt.Errorf("portal_group_not_found %w", err)
	}

	if err := s.DB.Delete(&portalGroup).Error; err != nil {
		return fmt.Errorf("failed_to_delete_portal_group: %w", err)
	}

	return s.WriteConfig(true)
}
