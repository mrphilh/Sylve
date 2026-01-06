package iscsi

import (
	"fmt"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
)

func (s *Service) GetLuns() ([]iscsiModels.Lun, error) {
	var luns []iscsiModels.Lun
	if err := s.DB.Find(&luns).Error; err != nil {
		return nil, fmt.Errorf("failed_to_get_luns: %w", err)
	}
	return luns, nil
}

func (s *Service) CreateLun(name string, path string, size string) error {
	if err := s.DB.Where("name = ?", name).First(&iscsiModels.Lun{}).Error; err == nil {
		return fmt.Errorf("lun_with_name_exists")
	}

	lun := iscsiModels.Lun{
		Name: name,
		Path: path,
		Size: size,
	}

	if err := s.DB.Create(&lun).Error; err != nil {
		return fmt.Errorf("failed_to_create_lun: %w", err)
	}

	return s.WriteConfig(true)
}

func (s *Service) UpdateLun(id uint, name string, path string, size string) error {
	var lun iscsiModels.Lun
	if err := s.DB.First(&lun, id).Error; err != nil {
		return fmt.Errorf("lun_not_found %w", err)
	}

	if name != lun.Name {
		var count int64
		if err := s.DB.Model(&iscsiModels.Lun{}).Where("name = ? AND id != ?", name, id).Count(&count).Error; err != nil {
			return fmt.Errorf("failed_to_check_name_confict: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("lun_with_name_exists")
		}

		tx := s.DB.Begin()

		lun.Name = name
		lun.Path = path
		lun.Size = size

		if err := tx.Save(&lun).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed_to_update_lun_fields: %w", err)
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed_to_commit_transaction: %w", err)
		}
	}
	return s.WriteConfig(true)
}

func (s *Service) DeleteLun(id uint) error {
	var lun iscsiModels.Lun
	if err := s.DB.First(&lun, id).Error; err != nil {
		return fmt.Errorf("lun_not_found %w", err)
	}

	if err := s.DB.Delete(&lun).Error; err != nil {
		return fmt.Errorf("failed_to_delete_lun: %w", err)
	}

	return s.WriteConfig(true)
}
