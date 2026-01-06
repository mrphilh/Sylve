package iscsi

import (
	"fmt"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
)

func (s *Service) GetAuthGroups() ([]iscsiModels.AuthGroup, error) {
	var authGroups []iscsiModels.AuthGroup
	if err := s.DB.Find(&authGroups).Error; err != nil {
		return nil, fmt.Errorf("failed_to_get_auth_groups: %w", err)
	}
	return authGroups, nil
}

func (s *Service) CreateAuthGroup(name string, authType string) error {
	if err := s.DB.Where("name = ?", name).First(&iscsiModels.AuthGroup{}).Error; err == nil {
		return fmt.Errorf("auth_group_with_name_exists")
	}

	authGroup := iscsiModels.AuthGroup{
		Name:     name,
		AuthType: authType,
	}

	if err := s.DB.Create(&authGroup).Error; err != nil {
		return fmt.Errorf("failed_to_create_auth_group: %w", err)
	}

	return s.WriteConfig(true)
}

func (s *Service) UpdateAuthGroup(id uint, name string, authType string) error {
	var authGroup iscsiModels.AuthGroup
	if err := s.DB.First(&authGroup, id).Error; err != nil {
		return fmt.Errorf("auth_group_not_found %w", err)
	}

	if name != authGroup.Name {
		var count int64
		if err := s.DB.Model(&iscsiModels.AuthGroup{}).Where("name = ? AND id != ?", name, id).Count(&count).Error; err != nil {
			return fmt.Errorf("failed_to_check_name_confict: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("auth_group_with_name_exists")
		}

		tx := s.DB.Begin()

		authGroup.Name = name
		authGroup.AuthType = authType

		if err := tx.Save(&authGroup).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed_to_update_auth_group_fields: %w", err)
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed_to_commit_transaction: %w", err)
		}
	}
	return s.WriteConfig(true)
}

func (s *Service) DeleteAuthGroup(id uint) error {
	var authGroup iscsiModels.AuthGroup
	if err := s.DB.First(&authGroup, id).Error; err != nil {
		return fmt.Errorf("auth_group_not_found %w", err)
	}

	if err := s.DB.Delete(&authGroup).Error; err != nil {
		return fmt.Errorf("failed_to_delete_auth_group: %w", err)
	}

	return s.WriteConfig(true)
}
