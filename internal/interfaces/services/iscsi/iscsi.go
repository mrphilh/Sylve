package iscsiServiceInterfaces

type IscsiServiceInterface interface {
	WriteConfig(reload bool) error
}
