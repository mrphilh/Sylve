package iscsiModels

import (
	"fmt"
	"strings"
)

/** TODO: Complete Authgroup configuration options
type Chap struct {
	User   string `json:"user"`
	Secret string `json:"secret"`
}

type ChapMutual struct {
	User         string `json:"user"`
	Secret       string `json:"secret"`
	MutualUser   string `json:"mutualUser"`
	MutualSecret string `json:"mutualSecret"`
}
*/

type AuthGroup struct {
	ID       int `json:"id" gorm:"primaryKey"`
	Name     string
	AuthType string `json:"authType" gorm:"default:'none'"`
	/*
		 TODO: Allow complete auth-group configuration options
			Chap            *Chap        `json:"chap"`
			ChapMutal       *[]ChapMutal `json:chapMutual`
			HostAddress     *string      `json:"hostAddress"`
			HostNqn         *string      `json:"hostNqn"`
			InitiatorName   *string      `json:"initiatorName"`
			InitiatorPortal *string      `json:"initiatorPortal"`
	*/
}

func (l *AuthGroup) AsUcl(indent int) string {
	builder := strings.Builder{}
	for i := 0; i < indent; i++ {
		builder.WriteString("\t")
	}
	tabs := builder.String()
	builder.Reset()
	builder.WriteString(fmt.Sprintf("%s%s {\n", tabs, l.Name))
	builder.WriteString(fmt.Sprintf("%s\tauth-type = %s\n", tabs, l.AuthType))
	builder.WriteString(fmt.Sprintf("%s}\n", tabs))
	return builder.String()
}
