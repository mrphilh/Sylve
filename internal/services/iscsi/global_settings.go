package iscsi

import (
	"fmt"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
)

func (s *Service) GetGlobalSettings() ([]iscsiModels.GlobalSetting, error) {
	var globalSettings []iscsiModels.GlobalSetting
	if err := s.DB.Find(&globalSettings).Error; err != nil {
		return nil, fmt.Errorf("failed_to_get_global_settings: %w", err)
	}
	return globalSettings, nil
}
