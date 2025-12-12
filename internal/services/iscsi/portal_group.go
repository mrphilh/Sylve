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
