package iscsi

import (
	"fmt"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
)

func (s *Service) GetTargets() ([]iscsiModels.Target, error) {
	var targets []iscsiModels.Target
	if err := s.DB.Preload("Target").Preload("PortalGroup").Preload("Lun").Find(&targets).Error; err != nil {
		return nil, fmt.Errorf("failed_to_get_targets: %w", err)
	}
	return targets, nil
}

func (s *Service) CreateTarget(name string, alias string, authGroupId int, portalGroupId int, luns []int) error {
	if err := s.DB.Where("name = ?", name).First(&iscsiModels.Target{}).Error; err == nil {
		return fmt.Errorf("target_with_name_exists")
	}

	var authGroup iscsiModels.AuthGroup
	if err := s.DB.Where("id = ?", authGroupId).First(&authGroup).Error; err != nil {
		return fmt.Errorf("create_target_auth_group_lookup_failed: %w", err)
	}

	var portalGroup iscsiModels.PortalGroup
	if err := s.DB.Where("id = ?", portalGroupId).First(&portalGroup).Error; err != nil {
		return fmt.Errorf("create_target_portal_group_lookup_failed: %w", err)
	}

	lunz := make([]iscsiModels.Lun, len(luns))
	for i := 0; i < len(luns); i++ {
		var lun iscsiModels.Lun
		if err := s.DB.Where("id = ?", luns[i]).First(&lun).Error; err != nil {
			return fmt.Errorf("create_target_lun_lookup_failed: %w", err)
		}
		lunz = append(lunz, lun)
	}

	target := iscsiModels.Target{
		Name:          name,
		Alias:         &alias,
		AuthGroupID:   &authGroupId,
		AuthGroup:     authGroup,
		PortalGroupID: &portalGroupId,
		PortalGroup:   portalGroup,
		Luns:          lunz,
	}

	if err := s.DB.Create(&target).Error; err != nil {
		return fmt.Errorf("failed_to_create_target: %w", err)
	}

	return s.WriteConfig(true)
}

func (s *Service) UpdateTarget(id uint, name string, authGroupId int) error {
	var target iscsiModels.Target
	if err := s.DB.First(&target, id).Error; err != nil {
		return fmt.Errorf("target_not_found %w", err)
	}

	if name != target.Name {
		var count int64
		if err := s.DB.Model(&iscsiModels.Target{}).Where("name = ? AND id != ?", name, id).Count(&count).Error; err != nil {
			return fmt.Errorf("failed_to_check_name_confict: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("target_with_name_exists")
		}

		tx := s.DB.Begin()

		target.Name = name
		var authGroup iscsiModels.AuthGroup
		if err := s.DB.Where("id = ?", authGroupId).First(&authGroup).Error; err != nil {
			return fmt.Errorf("create_target_auth_group_lookup_failed: %w", err)
		}
		target.AuthGroup = authGroup

		if err := tx.Save(&target).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed_to_update_target_fields: %w", err)
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed_to_commit_transaction: %w", err)
		}
	}
	return s.WriteConfig(true)
}

func (s *Service) DeleteTarget(id uint) error {
	var target iscsiModels.Target
	if err := s.DB.First(&target, id).Error; err != nil {
		return fmt.Errorf("target_not_found %w", err)
	}

	if err := s.DB.Delete(&target).Error; err != nil {
		return fmt.Errorf("failed_to_delete_target: %w", err)
	}

	return s.WriteConfig(true)
}
