package iscsiHandlers

import (
	"net/http"

	"github.com/alchemillahq/sylve/internal"
	iscsiModels "github.com/alchemillahq/sylve/internal/db/models/iscsi"

	"github.com/alchemillahq/sylve/internal/services/iscsi"
	"github.com/gin-gonic/gin"
)

// @Summary Get iSCSI Auth Group Configurations
// @Description Retrieve iSCSI Auth Group Configurations
// @Tags iSCSI
// @Accept json
// @Produce json
// @Success 200 {object} internal.APIResponse[[]iscsiModels.AuthGroup] "iSCSI Auth Group configurations"
// @Failure 500 {string} string "Internal server error"
// @Router /iscsi/authgroup [get]
func GetAuthGroups(iscsiService *iscsi.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth_groups, err := iscsiService.GetAuthGroups()
		if err != nil {
			c.JSON(http.StatusInternalServerError, internal.APIResponse[any]{
				Status:  "error",
				Message: "failed_to_get_iscsi_auth_groups",
				Error:   err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, internal.APIResponse[[]iscsiModels.AuthGroup]{
			Status:  "success",
			Message: "auth_groups_retrieved",
			Error:   "",
			Data:    auth_groups,
		})
	}
}
