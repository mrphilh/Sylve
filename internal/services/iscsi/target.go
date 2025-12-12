package iscsi

import (
	"fmt"

	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"
)

func (s *Service) GetTargets() ([]iscsiModels.Target, error) {
	var targets []iscsiModels.Target
	if err := s.DB.Preload("AuthGroup").Preload("PortalGroup").Preload("Lun").Find(&targets).Error; err != nil {
		return nil, fmt.Errorf("failed_to_get_targets: %w", err)
	}
	return targets, nil
}
