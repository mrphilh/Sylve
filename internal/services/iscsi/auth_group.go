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
