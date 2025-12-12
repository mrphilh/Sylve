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
