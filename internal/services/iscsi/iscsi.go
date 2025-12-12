package iscsi

import (
	iscsiServiceInterfaces "github.com/alchemillahq/sylve/internal/interfaces/services/iscsi"
	"gorm.io/gorm"
)

var _ iscsiServiceInterfaces.IscsiServiceInterface = (*Service)(nil)

type Service struct {
	DB *gorm.DB
}
